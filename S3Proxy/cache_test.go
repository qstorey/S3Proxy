package S3Proxy

import (
	"sync"
	"testing"
	"time"
)

func TestCacheBucketGet(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	//Needed for Options.BucketCacheTTL
	LoadDefaultOptions()
	// Lower the TTL for testing purposes
	Options.BucketCacheTTL = time.Duration(2 * time.Second)
	// Enter a bucket item into the cache
	CacheBucketSet("bucket_test_add", "eu-west-1")
	go func() {
		for i := 0; i < 10; i++ {
			CacheBucketSet("bucket_test_add", "eu-west-1")
		}
		wg.Done()
	}()
	// Retrieve the bucket item from the cache
	bucket := CacheBucketGet("bucket_test_add")
	// Confirm that the data is valid
	if bucket != nil {
		if bucket.Name != "bucket_test_add" {
			t.Error(
				"For", "bucket.Name",
				"Expected", "bucket_test_add",
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
			"Expected", "S3BucketCacheItem",
			"Got", nil,
		)
	}
	wg.Wait()
}

func TestCacheBucketExpire(t *testing.T) {
	//Needed for Options.BucketCacheTTL
	LoadDefaultOptions()
	// Lower the TTL for testing purposes
	Options.BucketCacheTTL = time.Duration(2 * time.Second)
	// Enter a bucket item into the cache
	CacheBucketSet("bucket_test_expire", "eu-west-1")
	// Sleep long enough for the entry to expire
	time.Sleep(Options.BucketCacheTTL + time.Second)
	// Retrieve the bucket item from the cache
	bucket := CacheBucketGet("bucket_test_expire")
	// Confirm that the entry expired
	if bucket != nil {
		t.Error(
			"For", "CacheBucketGet",
			"Expected", nil,
			"Got", bucket,
		)
	}

}

func TestCacheObjectGet(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	// Needed for Options.ObjectCacheTTL
	LoadDefaultOptions()
	// Lower the TTL for testing purposes
	Options.ObjectCacheTTL = time.Duration(2 * time.Second)
	// Enter an Object item into the cache
	CacheObjectSet("object_test_add", "bucket", "/file/path")
	go func() {
		for i := 0; i < 10; i++ {
			CacheObjectSet("object_test_add", "bucket", "/file/path")
		}
		wg.Done()
	}()
	// Retrieve the object item from the cache
	object := CacheObjectGet("object_test_add")
	if object != nil {
		if object.Key != "object_test_add" {
			t.Error(
				"For", "object.Key",
				"Expected", "object_test_add",
				"Got", object.Key,
			)
		}
		if object.Bucket != "bucket" {
			t.Error(
				"For", "object.Bucket",
				"Expected", "bucket",
				"Got", "object.Bucket",
			)
		}
		if object.FilePath != "/file/path" {
			t.Error(
				"For", "object.FilePath",
				"Expected", "bucket",
				"Got", "/file/path",
			)
		}
	} else {
		t.Error(
			"For", "object",
			"Expect", "SS3ObjectCacheItem",
			"Got", nil,
		)
	}
	wg.Wait()
}

func TestCacheObjectExpire(t *testing.T) {
	// Needed for Options.ObjectCacheTTL
	LoadDefaultOptions()
	// Lower the TTL for testing purposes
	Options.ObjectCacheTTL = time.Duration(2 * time.Second)
	// Enter an Object item into the cache
	CacheObjectSet("object_test_expire", "bucket", "/file/path")
	// Sleep long enough for the entry to expire
	time.Sleep(Options.ObjectCacheTTL + time.Second)
	object := CacheObjectGet("object_test_expire")
	// Confirm that the entry expired
	if object != nil {
		t.Error(
			"For", "CacheObjectGet",
			"Expected", "nil",
			"Got", object,
		)
	}
}
