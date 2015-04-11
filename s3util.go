package S3Proxy

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var s3bucket *s3.Bucket = nil

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
		"Key":          s3key.Key,
		"Size":         strconv.FormatInt(s3key.Size, 10),
		"LastModified": s3key.LastModified,
		"ETag":         s3key.ETag,
	}
	return resp, nil
}

func S3DownloadKey(key string) (string, error) {
	if s3bucket == nil {
		S3Connect()
	}
	body, err := s3bucket.Get(key)
	if err != nil {
		return "", err
	}
	filename := filepath.Clean(Options.CacheDir + key)
	// Create the subdirectories to match the key
	err = os.MkdirAll(filepath.Dir(filename), 0700)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		return "", err
	}
	return filename, nil
}
