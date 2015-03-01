package main

import (
    "fmt"
    "github.com/mitchellh/goamz/aws"
    "github.com/mitchellh/goamz/s3"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
)

const tmp_path string = "/tmp/s3http/"

func configure() {
    os.Mkdir(tmp_path, 0700)
}

func s3download(path string) {
    auth, err := aws.EnvAuth()
    if err != nil {
        log.Fatal(err)
    }
    client := s3.New(auth, aws.EUWest)
    bucket := client.Bucket(os.Getenv("AWS_BUCKET"))
    body, err := bucket.Get(path)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("about to write\n")
    filename := filepath.Clean(tmp_path + path)
    err = os.MkdirAll(filepath.Dir(filename), 0700)
    if err != nil {
        log.Fatal(err)
    }
    err = ioutil.WriteFile(filename, body, 0644)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Data written to path: %s", path)
}

func handler(w http.ResponseWriter, r *http.Request) {
    filename := tmp_path + r.URL.Path[1:]
    // If we don't have the file on disk, we need to download it
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        fmt.Printf("File isn't on disk so download it\n")
        s3download(r.URL.Path)
    }
    http.ServeFile(w, r, filename)
}

func main() {
    configure()
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

