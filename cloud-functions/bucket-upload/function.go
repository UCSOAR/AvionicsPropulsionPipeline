package function

import (
	"context"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	cloudutils "example.com/cloud-utils"
)

func BucketUpload(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	cloudutils.SetCorsHeaders(w, cloudutils.Cors{
		AllowOrigin:  "*",
		AllowMethods: []string{"POST", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	})

	// Handle request method
	switch r.Method {
	case "POST":
		log.Println("POST request received for /BucketUpload")
		break
	case "OPTIONS":
		log.Println("OPTIONS request received for /BucketUpload")
		w.WriteHeader(http.StatusOK)
		return
	default:
		log.Printf("Invalid method: %s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Limit size of uploaded file
	log.Printf("Content-Length: %d bytes", r.ContentLength)
	if r.ContentLength > cloudutils.MaxFileSize {
		log.Println("File size exceeds limit")
		http.Error(w, "File is too large", http.StatusBadRequest)
		return
	}

	// Parse form data
	log.Println("Parsing form data...")
	if err := r.ParseMultipartForm(cloudutils.MaxFileSize); err != nil {
		log.Printf("Error parsing form data: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Form data parsed successfully")

	// Get file from form data
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("Failed to retrieve file from form data: %v", err)
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	log.Printf("File received: %s, Size: %d bytes", header.Filename, header.Size)

	// Create a GCS client
	log.Println("Creating GCS client...")
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("Failed to create GCS client: %v", err)
		http.Error(w, "Failed to create GCS client", http.StatusInternalServerError)
		return
	}
	defer client.Close()
	log.Println("GCS client created successfully")

	// Upload file to GCS
	bucketName := cloudutils.BucketName
	log.Printf("Accessing bucket: %s", bucketName)
	bucket := client.Bucket(bucketName)
	obj := bucket.Object(header.Filename)

	// Check if the file already exists
	log.Printf("Checking if file already exists: %s", header.Filename)
	if _, err := obj.Attrs(ctx); err == nil {
		log.Printf("File %s already exists in bucket", header.Filename)
		http.Error(w, "File already exists", http.StatusConflict)
		return
	} else if err != storage.ErrObjectNotExist {
		log.Printf("Error checking file existence: %v", err)
		http.Error(w, "Failed to check file existence", http.StatusInternalServerError)
		return
	}

	// File does not exist, proceed to upload
	log.Printf("Uploading file: %s", header.Filename)
	writer := obj.NewWriter(ctx)
	if _, err := io.Copy(writer, file); err != nil {
		log.Printf("Failed to copy file to GCS: %v", err)
		http.Error(w, "Failed to copy file to GCS", http.StatusInternalServerError)
		return
	}

	// Close writer
	if err := writer.Close(); err != nil {
		log.Printf("Failed to close GCS writer: %v", err)
		http.Error(w, "Failed to close GCS writer", http.StatusInternalServerError)
		return
	}

	log.Printf("File %s uploaded successfully", header.Filename)
	w.WriteHeader(http.StatusOK)
}
