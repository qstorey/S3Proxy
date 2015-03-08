package main

import (
    "github.com/qstorey/S3Proxy"
    "net/http"
    "os"
)

// Constants
const tmp_path string = "/tmp/s3http/"

func configure() {
    os.Mkdir(tmp_path, 0700)
}

func main() {
    mux := S3Proxy.SetUpRoutes()
    http.Handle("/", mux)
    http.ListenAndServe(":8080", nil)

}
