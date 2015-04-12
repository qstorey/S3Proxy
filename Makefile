.PHONY: default build vet fmt

build: vet fmt
	go get -v
	cd S3Proxy && go build -v
	go build -v -o bin/S3Proxy

default: build

fmt:
	go fmt .
	cd S3Proxy && go fmt .

install: build
	go install -v

vet:
	go vet .
	cd S3Proxy && go vet .
