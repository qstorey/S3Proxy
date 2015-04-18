package S3Proxy

import (
	"os"
	"testing"
)

func TestConfigure(t *testing.T) {
	//needed for Options.CacheDir to be set
	LoadDefaultOptions()
	Configure()

	fileInfo, err := os.Stat(Options.CacheDir)

	if err != nil {
		if os.IsNotExist(err) {
			t.Error(
				"For", "Configure()",
				"expected", Options.CacheDir,
				"got", err,
			)
		} else {
			panic("Error performing os.Stat() call")
		}
	}

	//drwx------
	bits := 1<<(32-1) | int(0700)
	if fileInfo.Mode() != os.FileMode(bits) {
		t.Error(
			"For", "fileInfo.Mode()",
			"expected", os.FileMode(bits),
			"got", fileInfo.Mode(),
		)
	}

}
