package S3Proxy

import (
	"encoding/json"
	"time"
)

var s3Buckets = map[string]*s3BucketCacheItem{}

type s3BucketCacheItem struct {
	Name      string
	Location  string
	Timestamp time.Time
	TTL       time.Duration
}

func (b s3BucketCacheItem) String() string {
	tmp := &b
	tmp.TTL = b.TTL - time.Since(b.Timestamp)
	out, _ := json.Marshal(tmp)
	return string(out)
}

func CacheBucketGet(name string) *s3BucketCacheItem {
	bucket, hit := s3Buckets[name]
	if hit {
		// We need to check that the cache entry hasn't expired
		if time.Since(bucket.Timestamp) <= bucket.TTL {
			LogInfo("S3 Bucket Cache Hit - " + bucket.String())
			return bucket
		} else {
			delete(s3Buckets, name)
		}
	}
	// We didn't get a cache hit
	LogInfo("S3 Bucket Cache Miss - {Name:" + name + "}")
	return nil
}

func CacheBucketSet(name, location string) *s3BucketCacheItem {
	bucket := new(s3BucketCacheItem)
	bucket.Name = name
	bucket.Location = location
	bucket.Timestamp = time.Now()
	bucket.TTL = Options.BucketCacheTTL
	s3Buckets[name] = bucket
	LogInfo("S3 Bucket Cache Set - " + bucket.String())
	return bucket
}
