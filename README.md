S3Proxy
=========

An http proxy to S3 private buckets.

Table of contents
=================

- [S3Proxy](#s3proxy)
- [Table of contents](#table-of-contents)
- [Development](#development)

Development
===========

All the packages can be found in the root of the repo with its own Makefile to
build the packages.

The executable can be found in the S3Proxy directory. Running the Makefile in
this directory will produce an S3Proxy executable.

<pre>
./S3Proxy
+ S3Proxy
  |-- Makefile
  |-- S3Proxy.go
|-- handlers.go
|-- Makefile
|-- options.go
....
</pre>
