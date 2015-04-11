.PHONY: default build

default: build

build:
	go get -v
	go build -v
	cd S3Proxy && go install -v
