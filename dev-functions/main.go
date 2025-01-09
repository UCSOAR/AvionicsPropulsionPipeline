package main

import (
	"log"

	bucketUpload "example.com/cloud-functions/bucket-upload"
	getBucketUploads "example.com/cloud-functions/get-bucket-uploads"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

const devPort string = "8080"

func main() {
	// Register functions to test locally here
	funcframework.RegisterHTTPFunction("/BucketUpload", bucketUpload.BucketUpload)
	funcframework.RegisterHTTPFunction("/GetBucketUploads", getBucketUploads.GetBucketUploads)

	log.Printf("Development server listening on port %s", devPort)

	if err := funcframework.Start(devPort); err != nil {
		log.Fatal(err)
	}

	log.Printf("Server listening on port %s", devPort)
}
