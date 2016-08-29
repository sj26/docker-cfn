# CloudFormation Utils

Very basic implementation of [aws-cfn-bootstrap's](https://github.com/mesosphere/aws-cfn-bootstrap) cfn-signal for signally resource creating during CloudFormation creation and updates packages as a slim [Docker image](https://hub.docker.com/r/sj26/cfn/).

May be expanded to include wait conditions, cfn-init, etc., at some point if they seem useful.

## Usage

```
docker run -it --rm --network host sj26/cfn signal --stack "My Stack" --resource "Some Resource"
```
