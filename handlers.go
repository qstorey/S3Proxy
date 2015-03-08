package S3Proxy

import (
    "encoding/json"
    "os"
    "net/http"
)

// The index handler
func IndexHandler(w http.ResponseWriter, req *http.Request) {
    msgDict := map[string]string{"Name": "s3http", "Description": "An AWS S3 proxy server"}
    msgJson, err := json.Marshal(msgDict)
    if err != nil {
        LogFatal(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(msgJson)
}

// The status handler for determining the status of the server
func StatusHandler(w http.ResponseWriter, req *http.Request) {
    panic("Not Implemented")
}

// The default handler used for everything else
func DefaultHandler(w http.ResponseWriter, req *http.Request) {
    filename := tmp_path + req.URL.Path[1:]
    // If we don't have the file on disk, we need to download it
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        LogInfo("File isn't on disk so download it\n")
        S3download(req.URL.Path)
    }
    http.ServeFile(w, req, filename)
}
