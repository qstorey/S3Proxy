package S3Proxy

import "time"

type options struct {
	CacheDir       string
	BindAddress    string
	BucketCacheTTL time.Duration
	ObjectCacheTTL time.Duration
}

// Options is a globally available struct for storing runtime options
var Options = options{}

// LoadDefaultOptions is used to load in the default options into the globally
// available options struct. This is typically one of the first things called
// on start up. After this all options can be overridden
func LoadDefaultOptions() {
	Options.CacheDir = "/tmp/S3Proxy/"
	Options.BindAddress = ":9090"
	Options.BucketCacheTTL = time.Duration(1 * time.Hour)
	Options.ObjectCacheTTL = time.Duration(1 * time.Minute)
}
