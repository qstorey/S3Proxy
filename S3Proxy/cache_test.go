package S3Proxy

import (
	"testing"
	"time"
)

func TestCacheBucketGet(t *testing.T) {
	//Needed for Options.BucketCacheTTL
	LoadDefaultOptions()
	// Lower the TTL for testing purposes
	Options.BucketCacheTTL = time.Duration(2 * time.Second)
	// Enter a bucket item into the cache
	CacheBucketSet("test_add", "eu-west-1")
	// Retrieve the bucket item from the cache
	bucket := CacheBucketGet("test_add")
	// Confirm that the data is valid
	if bucket != nil {
		if bucket.Name != "test_add" {
			t.Error(
				"For", "bucket.Name",
				"Expected", "test_add",
				"Got", bucket.Name,
			)
		}
		if bucket.Location != "eu-west-1" {
			t.Error(
				"For", "bucket.Location",
				"Expected", "eu-west-1",
				"Got", bucket.Location,
			)
		}
	} else {
		t.Error(
			"For", "bucket",
			"Expected", "s3BucketCacheItem",
			"Got", nil,
		)
	}
}

func TestCacheBucketExpire(t *testing.T) {
	//Needed for Options.BucketCacheTTL
	LoadDefaultOptions()
	// Lower the TTL for testing purposes
	Options.BucketCacheTTL = time.Duration(2 * time.Second)
	// Enter a bucket item into the cache
	CacheBucketSet("test_expire", "eu-west-1")
	// Sleep long enough for the entry to expire
	time.Sleep(Options.BucketCacheTTL + time.Second)
	// Retrieve the bucket item from the cache
	bucket := CacheBucketGet("test_expire")
	// Confirm that the entry expired
	if bucket != nil {
		t.Error(
			"For", "CacheBucketGet",
			"Expected", nil,
			"Got", bucket,
		)
	}

}
