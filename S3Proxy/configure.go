package S3Proxy

import "os"

func Configure() {
	err := os.Mkdir(Options.CacheDir, 0700)
	if err != nil {
		if os.IsExist(err) {
			LogInfo("Directory " + Options.CacheDir + " already exists. Skipping mkdir.")
		} else {
			panic(err)
		}
	}
}
