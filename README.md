# S3Proxy

> WARNING: I created this repo for the intention of learning Go, use with caution.


S3Proxy is an http proxy to AWS S3 private buckets written in Go.

Using your AWS API credentials, S3Proxy allows you to make unauthenticated HTTP requests to download objects from your private S3 bucket.

This makes it easy for tools like curl to access private buckets:

```bash
curl http://localhost:9090/object/key
```

> WARNING: It's probably never a good idea to run this on a publicly accessible network/server, but rather private/local networks.

# Table of contents

- [S3Proxy](#s3proxy)
- [Table of contents](#table-of-contents)
- [Development](#development)
- [Usage](#usage)

# Development

Setup your development environment, see https://golang.org/doc/code.html. 
If you want to contribute to S3Proxy, clone this repo and run the ```make```

# Usage
```bash
export AWS_ACCESS_KEY=<your access key>
export AWS_SECRET_KEY=<you secret key>
export AWS_BUCKET=<the bucket you want to access>

$GOPATH/bin/S3Proxy
```
