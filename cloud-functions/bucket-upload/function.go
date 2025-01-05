package function

import (
	"context"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
)

const bucketName string = "gcf-v2-uploads-1026004530618.us-central1.cloudfunctions.appspot.com"
const maxFileSize int64 = 10 << 20 // 10MB

func BucketUpload(w http.ResponseWriter, r *http.Request) {
	SetCorsHeaders(w)

	switch r.Method {
	case "POST":
		{
			// Only allow POST requests
			break
		}
	case "OPTIONS":
		{
			w.WriteHeader(http.StatusOK)
			return
		}
	default:
		{
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}

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
	obj := bucket.Object(objectName)

	// Check if the file exists
	if _, err := obj.Attrs(ctx); err == nil {
		// File exists
		http.Error(w, "File already exists", http.StatusConflict)
		return
	} else if err != storage.ErrObjectNotExist {
		http.Error(w, "Failed to check file existence", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// File does not exist, proceed to upload
	writer := obj.NewWriter(ctx)

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

	w.WriteHeader(http.StatusOK)
}
