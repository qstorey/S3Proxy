package S3Proxy

import "time"

var s3Buckets = map[string]*s3BucketCacheItem{}
var s3Objects = map[string]*s3ObjectCacheItem{}

type s3BucketCacheItem struct {
	Name      string
	Location  string
	Timestamp time.Time
	TTL       time.Duration
}

func (b s3BucketCacheItem) String() string {
	return "{Name:" + b.Name +
		",Location:" + b.Location +
		",Timestamp:" + b.Timestamp.String() +
		",TTL:" + (b.TTL - time.Since(b.Timestamp)).String() +
		"}"
}

func CacheBucketGet(name string) *s3BucketCacheItem {
	bucket, hit := s3Buckets[name]
	if hit {
		// We need to check that the cache entry hasn't expired
		if time.Since(bucket.Timestamp) <= bucket.TTL {
			LogInfo("S3 Bucket Cache Hit - " + bucket.String())
			return bucket
		}
	}
	// We didn't get a cache hit
	LogInfo("S3 Bucket Cache Miss - {Name:" + name + "}")
	return nil
}

func CacheBucketSet(name string, location string) *s3BucketCacheItem {
	bucket := new(s3BucketCacheItem)
	bucket.Name = name
	bucket.Location = location
	bucket.Timestamp = time.Now()
	bucket.TTL = Options.BucketCacheTTL
	s3Buckets[name] = bucket
	LogInfo("S3 Bucket Cache Set - " + bucket.String())
	return bucket
}

type s3ObjectCacheItem struct {
	Bucket    string
	Key       string
	Timestamp time.Time
	TTL       time.Duration
}
