package S3Proxy

import "os"

func Configure() {
   os.Mkdir(Options.CacheDir, 0700)
}
