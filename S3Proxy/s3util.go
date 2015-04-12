package S3Proxy

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/s3"
	"io/ioutil"
	"os"
	"path/filepath"
)

type S3ProxyError struct {
	Code    int
	Message string
}

func handleError(e error) *S3ProxyError {
	err := new(S3ProxyError)
	if awserr := aws.Error(e); awserr != nil {
		err.Code = awserr.StatusCode
		err.Message = awserr.Code + ": " + awserr.Message
	} else if e != nil {
		// Not sure how to handle all errors will need to investigate this further.
		err.Code = 500
		err.Message = e.Error()
	}
	return err
}

func S3GetObject(bucket, key, region string) (string, *S3ProxyError) {
	svc := s3.New(&aws.Config{Region: region})
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	resp, err := svc.GetObject(params)
	if err != nil {
		LogError(err)
		return "", handleError(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		LogError(err)
		return "", handleError(err)
	}

	filename := filepath.Clean(Options.CacheDir + bucket + "/" + key)
	// Create the subdirectories to match the key
	err = os.MkdirAll(filepath.Dir(filename), 0700)
	if err != nil {
		LogError(err)
		return "", handleError(err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		LogError(err)
		return "", handleError(err)
	}
	return filename, nil
}
