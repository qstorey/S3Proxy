package S3Proxy

import (
    "fmt"
    "github.com/mitchellh/goamz/aws"
    "github.com/mitchellh/goamz/s3"
    "io/ioutil"
    "os"
    "path/filepath"

)

const tmp_path string = "/tmp/s3http/"

// S3 helpers
func S3download(path string) {
    auth, err := aws.EnvAuth()
    if err != nil {
        LogFatal(err)
    }
    client := s3.New(auth, aws.EUWest)
    bucket := client.Bucket(os.Getenv("AWS_BUCKET"))
    body, err := bucket.Get(path)
    if err != nil {
        LogFatal(err)
    }
    fmt.Println("about to write\n")
    filename := filepath.Clean(tmp_path + path)
    err = os.MkdirAll(filepath.Dir(filename), 0700)
    if err != nil {
        LogFatal(err)
    }
    err = ioutil.WriteFile(filename, body, 0644)
    if err != nil {
        LogFatal(err)
    }
    msg := "Data written to path: " + path
    LogInfo(msg)
}
