.PHONY: default build

default: build

build:
	go build -v
	cd S3Proxy && go build -v
