package main

import (
    "os"
    "net/http"
)

const tmp_path string = "/tmp/s3http/"

func configure() {
    os.Mkdir(tmp_path, 0700)
}

func handler(w http.ResponseWriter, r *http.Request) {
    filename := tmp_path + r.URL.Path[1:]
    http.ServeFile(w, r, filename)
}

func main() {
    configure()
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

