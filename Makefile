.PHONY: default build server

default: build

build:
	go get -v
	go build -v
	cd S3Proxy && go install -v

server: build
	${GOPATH}/bin/S3Proxy
