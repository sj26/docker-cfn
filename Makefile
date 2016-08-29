NAME=cfn
VERSION=$(shell git describe)
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

.PHONY: build
build: build/cfn build/ca-certificates.crt
	docker build --tag sj26/cfn:latest .

build/cfn: $(shell find . -type f -name "*.go")
	docker run --rm -it --volume ${PWD}:/go/src/github.com/sj26/docker-cfn --workdir /go/src/github.com/sj26/docker-cfn --env CGO_ENABLED=0 golang:1.7 go build ${LDFLAGS} -o build/${NAME} .

build/ca-certificates.crt:
	wget https://curl.haxx.se/ca/cacert.pem -O build/ca-certificates.crt

.PHONY: test
test:
	# XXX: Superficial test, but enough for now
	docker run --rm sj26/cfn --version

.PHONY: push
push:
	docker push sj26/cfn:latest

.PHONY: clean
clean:
	rm -f build/${NAME}
