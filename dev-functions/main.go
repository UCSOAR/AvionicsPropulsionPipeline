package main

import (
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	getStaticFireColumns "github.com/UCSOAR/AvionicsPropulsionPipeline/cloud-functions/get-static-fire-columns"
	getStaticFireMetadata "github.com/UCSOAR/AvionicsPropulsionPipeline/cloud-functions/get-static-fire-metadata"
	uploadStaticFire "github.com/UCSOAR/AvionicsPropulsionPipeline/cloud-functions/upload-static-fire"
)

const devPort string = "8080"

func main() {
	// Register functions to test locally here
	funcframework.RegisterHTTPFunction("/GetStaticFireMetadata", getStaticFireMetadata.GetStaticFires)
	funcframework.RegisterHTTPFunction("/GetStaticFireColumns", getStaticFireColumns.GetStaticFireColumns)
	funcframework.RegisterHTTPFunction("/UploadStaticFire", uploadStaticFire.UploadStaticFire)

	log.Printf("Development server listening on port %s", devPort)

	if err := funcframework.Start(devPort); err != nil {
		log.Fatal(err)
	}

	log.Printf("Server listening on port %s", devPort)
}
