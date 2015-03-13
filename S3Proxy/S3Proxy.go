package main

import (
    "github.com/qstorey/S3Proxy"
    "net/http"
    "os"
)

func configure(options *S3Proxy.Options) {
    os.Mkdir(options.CacheDir, 0700)
}

func main() {
    options := new(S3Proxy.Options)
    // Load the default options
    options.LoadDefaultOptions()
    // Run the startup configuration
    configure(options)
    // Set up the routing
    mux := S3Proxy.SetUpRoutes()
    http.Handle("/", mux)
    // Start the HTTP serer
    http.ListenAndServe(":8080", nil)
}
