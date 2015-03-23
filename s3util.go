package S3Proxy

import (
    "github.com/mitchellh/goamz/aws"
    "github.com/mitchellh/goamz/s3"
    "os"

)

var s3bucket *s3.Bucket
// S3 helpers

func S3Connect() {
  auth, err := aws.EnvAuth()
  if err != nil {
      LogFatal(err)
  }
  s3client := s3.New(auth, aws.EUWest)
  s3bucket = s3client.Bucket(os.Getenv("AWS_BUCKET"))
}
