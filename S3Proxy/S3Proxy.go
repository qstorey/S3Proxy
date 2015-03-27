package main

import (
    "github.com/qstorey/S3Proxy"
    "net/http"
    "os"
)

func configure() {
    os.Mkdir(S3Proxy.Options.CacheDir, 0700)
}

func main() {
    // Load the default options
    S3Proxy.LoadDefaultOptions()
    // Run the startup configuration
    configure()
    // Connect to S3
    S3Proxy.S3Connect()
    // Set up the routing
    mux := S3Proxy.SetUpRoutes()
    http.Handle("/", mux)
    // Start the HTTP serer
    http.ListenAndServe(S3Proxy.Options.BindAddress, nil)
}
