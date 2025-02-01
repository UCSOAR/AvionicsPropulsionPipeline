package main

import (
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	getStaticFires "github.com/UCSOAR/AvionicsPropulsionPipeline/cloud-functions/get-static-fires"
	uploadStaticFire "github.com/UCSOAR/AvionicsPropulsionPipeline/cloud-functions/upload-static-fire"
)

const devPort string = "8080"

func main() {
	// Register functions to test locally here
	funcframework.RegisterHTTPFunction("/UploadStaticFire", uploadStaticFire.UploadStaticFire)
	funcframework.RegisterHTTPFunction("/GetStaticFires", getStaticFires.GetStaticFires)

	log.Printf("Development server listening on port %s", devPort)

	if err := funcframework.Start(devPort); err != nil {
		log.Fatal(err)
	}

	log.Printf("Server listening on port %s", devPort)
}
