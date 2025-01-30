package main

import (
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	bucketUpload "github.com/UCSOAR/AvionicsPropulsionPipeline/cloud-functions/bucket-upload"
	downloadBucketObject "github.com/UCSOAR/AvionicsPropulsionPipeline/cloud-functions/download-bucket-object"
	getBucketUploads "github.com/UCSOAR/AvionicsPropulsionPipeline/cloud-functions/get-bucket-uploads"
)

const devPort string = "8080"

func main() {
	// Register functions to test locally here
	funcframework.RegisterHTTPFunction("/BucketUpload", bucketUpload.BucketUpload)
	funcframework.RegisterHTTPFunction("/GetBucketUploads", getBucketUploads.GetBucketUploads)
	funcframework.RegisterHTTPFunction("/DownloadBucketObject", downloadBucketObject.DownloadBucketObject)

	log.Printf("Development server listening on port %s", devPort)

	if err := funcframework.Start(devPort); err != nil {
		log.Fatal(err)
	}

	log.Printf("Server listening on port %s", devPort)
}
