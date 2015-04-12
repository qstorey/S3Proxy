package S3Proxy

import (
	"github.com/drone/routes"
	"net/http"
	"strings"
)

// The index handler
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	msgDict := map[string]string{"Name": "S3Proxy", "Description": "An AWS S3 proxy server"}
	routes.ServeJson(w, &msgDict)
	return
}

// The status handler for determining the status of the server
func StatusHandler(w http.ResponseWriter, req *http.Request) {
	panic("Not Implemented")
}

// The default handler used for everything else
func DefaultHandler(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	parts := strings.SplitN(req.URL.Path[1:], "/", 2)
	bucket, key := parts[0], parts[1]
	obj, err := S3GetObject(bucket, key, params.Get("aws_region"))
	if err != nil {
		http.Error(w, err.Message, err.Code)
	}
	http.ServeFile(w, req, obj)
	return
}
