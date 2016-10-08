package S3Proxy

import (
	"encoding/json"
	"sync"
	"time"
)

var mutexB = sync.Mutex{}
var mutexO = sync.Mutex{}
var s3Buckets = map[string]*S3BucketCacheItem{}
var s3Objects = map[string]*S3ObjectCacheItem{}

type s3CacheItem struct {
	Timestamp time.Time
	TTL       time.Duration
}

func (c s3CacheItem) String() string {
	tmp := &c
	tmp.TTL = c.TimeLeft()
	out, _ := json.Marshal(tmp)
	return string(out)
}

// TimeLeft returns the time duration before the CacheItem expires
func (c s3CacheItem) TimeLeft() time.Duration {
	return c.TTL - time.Since(c.Timestamp)
}

// S3BucketCacheItem is a cache entry representing bucket info
type S3BucketCacheItem struct {
	CacheItem s3CacheItem
	Name      string
	Location  string
}

func (b S3BucketCacheItem) String() string {
	out, _ := json.Marshal(b)
	return string(out)
}

// S3ObjectCacheItem represents a cache entry for meta data for an S3 object
// and not the S3Object itself. The idea is to reduce S3 API calls for Object
// information
type S3ObjectCacheItem struct {
	CacheItem s3CacheItem
	Key       string
	Bucket    string
	FilePath  string
}

func (o S3ObjectCacheItem) String() string {
	out, _ := json.Marshal(o)
	return string(out)
}

// CacheBucketGet checks if the bucket information is in the local cache
func CacheBucketGet(name string) *S3BucketCacheItem {
	mutexB.Lock()
	defer mutexB.Unlock()
	bucket, hit := s3Buckets[name]
	if hit {
		// We need to check that the cache entry hasn't expired
		if time.Since(bucket.CacheItem.Timestamp) <= bucket.CacheItem.TTL {
			LogInfo("S3 Bucket Cache Hit - " + bucket.String())
			return bucket
		}
		// The cache entry has expired so remove it
		delete(s3Buckets, name)
	}
	// We didn't get a cache hit
	LogInfo("S3 Bucket Cache Miss - {\"Name\":\"" + name + "\"}")
	return nil
}

// CacheBucketSet enters a bucket's info into the cache
func CacheBucketSet(name, location string) *S3BucketCacheItem {
	bucket := S3BucketCacheItem{
		CacheItem: s3CacheItem{Timestamp: time.Now(), TTL: Options.BucketCacheTTL},
		Name:      name,
		Location:  location,
	}
	mutexB.Lock()
	defer mutexB.Unlock()
	s3Buckets[name] = &bucket
	LogInfo("S3 Bucket Cache Set - " + bucket.String())
	return &bucket
}

// CacheObjectGet checks if the S3Object information is in the local cache
func CacheObjectGet(key string) *S3ObjectCacheItem {
	mutexO.Lock()
	defer mutexO.Unlock()
	object, hit := s3Objects[key]
	if hit {
		if time.Since(object.CacheItem.Timestamp) <= object.CacheItem.TTL {
			LogInfo("S3 Object Cache Hit - " + object.String())
			return object
		}
		// The cache entry has expired so remove it
		delete(s3Objects, key)
	}
	LogInfo("S3 Object Cache Miss - {\"Key\":\"" + key + "\"}")
	return nil
}

// CacheObjectSet enters a S3 Object's info into the cache
func CacheObjectSet(key, bucket, filepath string) *S3ObjectCacheItem {
	object := S3ObjectCacheItem{
		CacheItem: s3CacheItem{Timestamp: time.Now(), TTL: Options.ObjectCacheTTL},
		Key:       key,
		Bucket:    bucket,
		FilePath:  filepath,
	}
	mutexO.Lock()
	defer mutexO.Unlock()
	s3Objects[key] = &object
	LogInfo("S3 Object Cache Set - " + object.String())
	return &object
}
