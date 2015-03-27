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
    routes.ServeJson(w, &keyMap)
    return
    // // Check if the key is still valid on S3

    // filename := Options.CacheDir + req.URL.Path[1:]
    // // If we don't have the file on disk, we need to download it
    // if _, err := os.Stat(filename); os.IsNotExist(err) {
    //     LogInfo("File isn't on disk so download it\n")
    //     // S3download(req.URL.Path)
    // }
    // http.ServeFile(w, req, filename)
}
