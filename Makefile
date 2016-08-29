NAME=cfn
VERSION=$(git describe)
LDFLAGS=-ldflags "-X github.com/sj26/docker-cfn/main.Version=${VERSION}"
CGO_ENABLED=0

build/cfn:
	mkdir -p build
	go build ${LDFLAGS} -o build/${NAME} *.go

build/ca-certificates.crt:
	mkdir -p build
	wget https://curl.haxx.se/ca/cacert.pem -O build/ca-certificates.crt

.PHONY: install
install:
	go install ${LDFLAGS} .

.PHONY: clean
clean:
	rm -f build/${NAME}
