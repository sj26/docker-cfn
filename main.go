package main

import (
  "os"

  "github.com/urfave/cli"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/ec2metadata"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/cloudformation"
)

// Populated by ldflags
var Version string

func main() {
  app := cli.NewApp()
  app.Name = "cfn"
  app.Version = Version
  app.Flags = []cli.Flag{
    cli.StringFlag{
      Name: "region",
      EnvVar: "DEFAULT_AWS_REGION,AWS_REGION",
    },
  }
  app.Commands = []cli.Command{
    cli.Command{
      Name: "signal",
      Flags: []cli.Flag{
        cli.StringFlag{
          Name: "stack",
        },
        cli.StringFlag{
          Name: "resource",
        },
        cli.StringFlag{
          Name: "id, i",
        },
        cli.BoolFlag{
          Name: "success",
        },
        cli.IntFlag{
          Name: "exit-code",
        },
      },
      Action: signalCommand,
    },
  }
  app.Run(os.Args)
}

func signalCommand(c *cli.Context) (err error) {
  sess := session.New()
  ec2m := ec2metadata.New(sess)
  ec2q := false
  var ec2i ec2metadata.EC2InstanceIdentityDocument

  var region string
  if c.GlobalIsSet("region") {
    region = c.GlobalString("region")
  } else {
    ec2i, err = ec2m.GetInstanceIdentityDocument()
    if err != nil {
      return err
    } else {
      ec2q = true
    }
    region = ec2i.AvailabilityZone[:len(ec2i.AvailabilityZone)-1]
  }

  config := &aws.Config{Region: aws.String(region)}
  cfn := cloudformation.New(sess, config)

  stack := c.String("stack")
  if stack == "" {
    return cli.NewExitError("Stack name is required", 1)
  }

  resource := c.String("resource")
  if resource == "" {
    return cli.NewExitError("Logical resource ID is required", 1)
  }

  var id string
  if c.IsSet("id") {
    id = c.String("id")
  } else {
    if !ec2q {
      ec2i, err = ec2m.GetInstanceIdentityDocument()
      if err != nil {
        return err
      } else {
        ec2q = true
      }
    }
    id = ec2i.InstanceID
  }

  success := true

  if c.IsSet("success") {
    success = success && c.Bool("success")
  }

  if c.IsSet("exit-code") {
    success = success && c.Int("exit-code") == 0
  }

  var status string
  if success {
    status = cloudformation.ResourceSignalStatusSuccess
  } else {
    status = cloudformation.ResourceSignalStatusFailure
  }

  _, err = cfn.SignalResource(&cloudformation.SignalResourceInput{
    StackName: aws.String(stack),
    LogicalResourceId: aws.String(resource),
    UniqueId: aws.String(id),
    Status: aws.String(status),
  })
  if err != nil {
    return err
  }

  return nil
}
