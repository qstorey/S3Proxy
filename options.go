package S3Proxy

type Options struct {
    CacheDir    string
}

func (opt *Options) LoadDefaultOptions() {
    opt.CacheDir = "/tmp/S3Proxy"
}
