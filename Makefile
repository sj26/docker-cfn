NAME=cfn
VERSION=$(git describe)
LDFLAGS=-ldflags "-X github.com/sj26/docker-cfn/main.Version=${VERSION}"

build/cfn:
	mkdir -p build
	docker run --rm -it --volume ${PWD}:/go/src/github.com/sj26/docker-cfn --workdir /go/src/github.com/sj26/docker-cfn --env CGO_ENABLED=0 golang:1.7 go build ${LDFLAGS} -o build/${NAME} .

build/ca-certificates.crt:
	mkdir -p build
	wget https://curl.haxx.se/ca/cacert.pem -O build/ca-certificates.crt

.PHONY: install
install:
	go install ${LDFLAGS} .

.PHONY: clean
clean:
	rm -f build/${NAME}
