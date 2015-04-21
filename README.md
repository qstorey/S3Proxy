# S3Proxy

S3Proxy is an http proxy to AWS S3 private buckets written in Go.

Using your AWS API credentials, S3Proxy allows you to make unauthenticated HTTP requests to download objects from your private S3 bucket.

This makes it easy for tools like curl to access private buckets:

Eg. Assuming you have a S3 bucket named `bucket1` and it has the following 
structure `folder1/folder2/item.txt`. 

```bash
curl http://localhost:9090/bucket1/folder1/folder2/item.txt
```

> WARNING: It's probably never a good idea to run this on a publicly accessible network/server, but rather private/local networks.

# Table of contents

- [S3Proxy](#s3proxy)
- [Table of contents](#table-of-contents)
- [Prerequisites](#Prerequisites)
- [Development](#development)
- [Usage](#usage)

# Prerequisites

Setup your development environment, see https://golang.org/doc/code.html.

Running the build script requires the ```go vet``` and ```go cover``` tool, which is obtained as follows:
```bash
go get golang.org/x/tools/cmd/vet
go get golang.org/x/tools/cmd/cover
```

# Development

If you want to make modifications to S3Proxy, clone this repo, ensure that all the prerequisites are met, make your changes and run the build script ```make```.

Pull requests are more than welcome if you would like to contribute back to S3Proxy.

# Usage

It is best to use this with an AWS IAM user with a minimal access policy. See
the following example policy. This provides s3 ReadOnly access to all your
buckets. Just add a Deny statement for each bucket you want to keep private.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:Get*",
        "s3:List*"
      ],
      "Resource": "*"
    }
  ]
}
```

```bash
export AWS_ACCESS_KEY=<your access key>
export AWS_SECRET_KEY=<you secret key>

./bin/S3Proxy
```
