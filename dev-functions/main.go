package main

import (
	"log"

	bucketUpload "example.com/cloud-functions/bucket-upload"
	downloadBucketObject "example.com/cloud-functions/download-bucket-object"
	getBucketUploads "example.com/cloud-functions/get-bucket-uploads"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

const devPort string = "8080"

func main() {
	// Register functions
	funcframework.RegisterHTTPFunction("/BucketUpload", bucketUpload.BucketUpload)
	log.Println("Registered HTTP function: /BucketUpload")

	funcframework.RegisterHTTPFunction("/GetBucketUploads", getBucketUploads.GetBucketUploads)
	log.Println("Registered HTTP function: /GetBucketUploads")

	funcframework.RegisterHTTPFunction("/DownloadBucketObject", downloadBucketObject.DownloadBucketObject)
	log.Println("Registered HTTP function: /DownloadBucketObject")

	// Start the development server
	log.Printf("Development server listening on port %s", devPort)
	if err := funcframework.Start(devPort); err != nil {
		log.Fatal(err)
	}

	log.Printf("Server listening on port %s", devPort)
}
