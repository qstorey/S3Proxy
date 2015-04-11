package S3Proxy

import "log"

// Logging
func LogInfo(s string) {
	log.Print(s)
}

func LogFatal(e error) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Fatal(e)
}
