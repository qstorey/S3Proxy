package S3Proxy

import (
	"github.com/drone/routes"
	"net/http"
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
	keyMap, err := S3ValidateKey(req.URL.Path[1:])
	if err != nil {
		http.Error(w, "404: Not Found", 404)
		return
	}
	filename, err := S3DownloadKey(keyMap["Key"])
	if err != nil {
		http.Error(w, "Server Exception", 500)
		return
	}
	http.ServeFile(w, req, filename)
	return
}
