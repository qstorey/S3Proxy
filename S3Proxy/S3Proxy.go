package main

import (
    "github.com/qstorey/S3Proxy"
    "net/http"
    "os"
)

var options S3Proxy.Options

func configure() {
    os.Mkdir(options.CacheDir, 0700)
}

func main() {
    // Load the default options
    options.LoadDefaultOptions()
    // Run the startup configuration
    configure()
    // Set up the routing
    mux := S3Proxy.SetUpRoutes()
    http.Handle("/", mux)
    // Start the HTTP server
    http.ListenAndServe(":8080", nil)
}
