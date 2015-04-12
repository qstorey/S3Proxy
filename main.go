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
	// Connect to S3
	//S3Proxy.S3Connect()
	// Set up the routing
	mux := S3Proxy.SetUpRoutes()
	http.Handle("/", mux)
	// Start the HTTP serer
	http.ListenAndServe(S3Proxy.Options.BindAddress, nil)
}
