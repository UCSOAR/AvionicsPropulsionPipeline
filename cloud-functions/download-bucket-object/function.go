package function

import (
	"context"
	"net/http"

	"cloud.google.com/go/storage"
	cloudUtils "example.com/cloud-utils"
)

func DownloadBucketObject(w http.ResponseWriter, r *http.Request) {
	cloudUtils.SetCorsHeaders(w, cloudUtils.Cors{
		AllowOrigin:  "*",
		AllowMethods: []string{"GET", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Content-Disposition", "Authorization"},
	})

	switch r.Method {
	case "GET":
		{
			// Only allow GET requests
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

	// Parse file name from URL
	filename := r.URL.Query().Get("file")

	if filename == "" {
		http.Error(w, "No file specified", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// Create a GCS client
	client, err := storage.NewClient(ctx)

	if err != nil {
		http.Error(w, "Failed to create GCS client", http.StatusInternalServerError)
		return
	}

	defer client.Close()

	// Get a reader for the object
	bucket := client.Bucket(cloudUtils.BucketName)
	reader, err := bucket.Object(filename).NewReader(ctx)

	if err != nil {
		http.Error(w, "Failed to create reader for file", http.StatusInternalServerError)
		return
	}

	defer reader.Close()

	// Set response headers
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)

	// Copy object to response
	if _, err := reader.WriteTo(w); err != nil {
		http.Error(w, "Failed to copy file object to response", http.StatusInternalServerError)
		return
	}
}
