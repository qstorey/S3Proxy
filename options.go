package S3Proxy

type options struct {
    CacheDir    string
}

var Options = options{}

func LoadDefaultOptions() {
    Options.CacheDir = "/tmp/S3Proxy/"
}
