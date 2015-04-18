package S3Proxy

import "testing"

var defaultCacheDir = "/tmp/S3Proxy/"
var defaultBindAddress = ":9090"

func TestLoadDefaultOptions(t *testing.T) {
	LoadDefaultOptions()
	if Options.CacheDir != defaultCacheDir {
		t.Error(
			"For", "Options.CacheDir",
			"expected", defaultCacheDir,
			"got", Options.CacheDir,
		)
	}
	if Options.BindAddress != defaultBindAddress {
		t.Error(
			"For", "Options.BindAddress",
			"expected", defaultBindAddress,
			"got", Options.BindAddress,
		)
	}
}
