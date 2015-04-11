.PHONY: default build

default: build

build:
	go get -v
	cd S3Proxy && go build -v
	go build -v -o bin/S3Proxy
