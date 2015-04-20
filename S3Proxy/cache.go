package S3Proxy

import "time"

type s3BucketCacheItem struct {
	Name      string
	Location  string
	Timestamp time.Time
	TTL       time.Duration
}

type s3ObjectCacheItem struct {
	Bucket    string
	Key       string
	Timestamp time.Time
	TTL       time.Duration
}

var s3Buckets = map[string]*s3BucketCacheItem{}
var s3Objects = map[string]*s3ObjectCacheItem{}

func CacheBucketAdd(name string, location string) *s3BucketCacheItem {
	bucket := new(s3BucketCacheItem)
	bucket.Name = name
	bucket.Location = location
	bucket.Timestamp = time.Now()
	bucket.TTL = Options.BucketCacheTTL
	s3Buckets[name] = bucket
	return bucket
}

func CacheBucketGet(name string) *s3BucketCacheItem {
	bucket, hit := s3Buckets[name]
	if !hit {
		// We didn't get a cache hit
		return nil
	}
	// We need to check that the cache entry hasn't expired
	if time.Since(bucket.Timestamp) <= bucket.TTL {
		return bucket
	}
	return nil
}
