package S3Proxy

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3ProxyError struct {
	Code    int
	Message string
}

func handleError(e error) *S3ProxyError {
	err := new(S3ProxyError)
	if awsErr, ok := e.(awserr.RequestFailure); ok {
		err.Code = awsErr.StatusCode()
		err.Message = awsErr.Code() + ": " + awsErr.Message()
	} else if e != nil {
		// Not sure how to handle all errors will need to investigate this further.
		err.Code = 500
		err.Message = e.Error()
	}
	return err
}

// S3GetBucketLocation returns the AWS region of the S3 Bucket
func S3GetBucketLocation(bucket string) (string, *S3ProxyError) {
	// Check if we have the bucket cached
	bucketCacheItem := CacheBucketGet(bucket)
	if bucketCacheItem != nil {
		return bucketCacheItem.Location, nil
	}
	// Strange behaviour when hitting s3.amazonaws.com. Some regions work fine
	// other return AuthorizationMalformedHeader. Specifying a region other than
	// us-east-1 always works.
	svc := s3.New(session.New(&aws.Config{Region: aws.String("eu-west-1")}))
	params := &s3.GetBucketLocationInput{
		Bucket: aws.String(bucket),
	}
	resp, err := svc.GetBucketLocation(params)
	if err != nil {
		LogError(err)
		return "", handleError(err)
	}

	// API returns the empty response when bucket location is US Standard
	awsRegion := "us-east-1"
	if resp.LocationConstraint != nil {
		awsRegion = *resp.LocationConstraint
	}
	CacheBucketSet(bucket, awsRegion)
	return awsRegion, nil
}

// S3GetObject returns path to the file on disk where the S3 object has been
// downloaded to
func S3GetObject(bucket, key, region string) (string, *S3ProxyError) {
	// Check if the item
	objectCacheItem := CacheObjectGet(key)
	if objectCacheItem != nil {
		return objectCacheItem.FilePath, nil
	}
	svc := s3.New(session.New(&aws.Config{Region: aws.String(region)}))
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
	CacheObjectSet(key, bucket, filename)
	return filename, nil
}
