package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

const bucketName string = "gcf-v2-uploads-1026004530618.us-central1.cloudfunctions.appspot.com"
const maxFileSize int64 = 10 << 20 // 10MB
const devPort string = "8080"

func BucketUpload(w http.ResponseWriter, r *http.Request) {
	// Limit size of uploaded file
	if r.ContentLength > maxFileSize {
		http.Error(w, "File is too large", http.StatusBadRequest)
		return
	}

	// Parse form data
	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get file from form data
	file, header, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Create a GCS client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		http.Error(w, "Failed to create GCS client", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	defer client.Close()

	// Upload file to GCS
	objectName := header.Filename
	bucket := client.Bucket(bucketName)
	writer := bucket.Object(objectName).NewWriter(ctx)

	// Copy file to GCS
	if _, err := io.Copy(writer, file); err != nil {
		http.Error(w, "Failed to copy file to GCS", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Close writer
	if err := writer.Close(); err != nil {
		http.Error(w, "Failed to close GCS writer", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "File uploaded to %s", objectName)

	// Respond with success
	w.WriteHeader(http.StatusOK)
}

func main() {
	funcframework.RegisterHTTPFunction("/upload", BucketUpload)

	log.Printf("Listening on port %s", devPort)
	log.Fatal(funcframework.Start(devPort))
}
