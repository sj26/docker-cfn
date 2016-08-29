FROM scratch
MAINTAINER sj26@sj26.com
ADD build/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ADD build/cfn /cfn
ENTRYPOINT ["/cfn"]
