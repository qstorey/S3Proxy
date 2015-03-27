package S3Proxy

import (
    "github.com/mitchellh/goamz/aws"
    "github.com/mitchellh/goamz/s3"
    "os"
    "strconv"
)

var s3bucket *s3.Bucket = nil
// S3 helpers

func S3Connect() {
  auth, err := aws.EnvAuth()
  if err != nil {
      LogFatal(err)
  }
  LogInfo("Authenticated to AWS")
  s3client := s3.New(auth, aws.EUWest)
  s3bucket = s3client.Bucket(os.Getenv("AWS_BUCKET"))
}

func S3ValidateKey(key string) (map[string]string, error) {
    if s3bucket == nil {
        S3Connect()
    }
    s3key, err := s3bucket.GetKey(key)
    if err != nil {
      return nil, err
    }

    resp := map[string]string{
      "Key": s3key.Key,
      "LastModified": s3key.LastModified,
      "Size": strconv.FormatInt(s3key.Size, 10),
    }
    return resp, nil
}
