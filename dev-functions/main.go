package main

import (
	"log"

	bucketUpload "example.com/cloud-functions/bucket-upload"
	helloWorld "example.com/cloud-functions/hello-world"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

const devPort uint32 = 8080

func main() {
	// Register functions to test locally here
	funcframework.RegisterHTTPFunction("/HelloWorld", helloWorld.HelloWorld)
	funcframework.RegisterHTTPFunction("/BucketUpload", bucketUpload.BucketUpload)

	log.Printf("Development server listening on port %d", devPort)

	if err := funcframework.Start(string(rune(devPort))); err != nil {
		log.Fatal(err)
	}

	log.Printf("Server listening on port %d", devPort)
}
