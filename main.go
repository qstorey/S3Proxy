package main

import (
	"github.com/qstorey/S3Proxy/S3Proxy"
	"net/http"
)

func main() {
	// Load the default options
	S3Proxy.LoadDefaultOptions()
	// Run the startup configuration
	S3Proxy.Configure()
	// Set up the routing
	mux := S3Proxy.SetUpRoutes()
	http.Handle("/", mux)
	// Start the HTTP serer
	S3Proxy.LogInfo("Listening on port: " + S3Proxy.Options.BindAddress)
	http.ListenAndServe(S3Proxy.Options.BindAddress, nil)
}
