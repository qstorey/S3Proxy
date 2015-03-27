package S3Proxy

type options struct {
    CacheDir    string
    BindAddress string
}

var Options = options{}

func LoadDefaultOptions() {
    Options.CacheDir = "/tmp/S3Proxy/"
    Options.BindAddress = ":9090"
}
