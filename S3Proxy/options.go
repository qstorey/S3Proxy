package S3Proxy

import "time"

type options struct {
	CacheDir       string
	BindAddress    string
	BucketCacheTTL time.Duration
}

var Options = options{}

func LoadDefaultOptions() {
	Options.CacheDir = "/tmp/S3Proxy/"
	Options.BindAddress = ":9090"
	Options.BucketCacheTTL = time.Duration(1 * time.Hour)
}

// Options to be loaded in for unit tests
func LoadTestOptions() {
	Options.CacheDir = "/tmp/S3Proxy/"
	Options.BindAddress = ":9090"
	Options.BucketCacheTTL = time.Duration(2 * time.Second)
}
