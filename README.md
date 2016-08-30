# CloudFormation Utils

Very basic implementation of [CloudFormation Helper Scripts'](http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-helper-scripts-reference.html) cfn-signal for signally resource creating during CloudFormation creation and updates packages as a slim [Docker image](https://hub.docker.com/r/sj26/cfn/).

May be expanded to include wait conditions, cfn-init, etc., at some point if they seem useful.

## Usage

```
docker run -it --rm --network host sj26/cfn signal --stack "My Stack" --resource "Some Resource"
```

The options available are designed to match [cfn-signal](http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-signal.html).

### cfn --help

```
NAME:
   cfn - A new cli application

USAGE:
   cfn [global options] command [command options] [arguments...]

VERSION:
   v1.0.0

COMMANDS:
     signal
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --region value   [$DEFAULT_AWS_REGION, $AWS_REGION]
   --help, -h      show help
   --version, -v   print the version
```

### cfn signal --help

```
NAME:
   cfn signal -

USAGE:
   cfn signal [command options] [arguments...]

OPTIONS:
   --stack value
   --resource value
   --id value, -i value
   --success
   --exit-code value     (default: 0)
```
