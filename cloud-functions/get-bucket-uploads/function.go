package function

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	cloudutils "example.com/cloud-utils"
	"google.golang.org/api/iterator"
)

func GetBucketUploads(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	cloudutils.SetCorsHeaders(w, cloudutils.Cors{
		AllowOrigin:  "*",
		AllowMethods: []string{"GET", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	})

	switch r.Method {
	case "GET":
		{
			log.Println("GET request received for /GetBucketUploads")
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

	ctx := context.Background()

	// Create a GCS client
	log.Println("Creating GCS client...")
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("Failed to create GCS client: %v", err)
		http.Error(w, "Failed to create GCS client", http.StatusInternalServerError)
		return
	}
	defer client.Close()
	log.Println("GCS client created successfully")

	// Access the specified bucket
	log.Printf("Accessing bucket: %s", cloudutils.BucketName)
	bucket := client.Bucket(cloudutils.BucketName)
	obj_it := bucket.Objects(ctx, nil)

	var objects []cloudutils.BucketObject

	// Fetch objects in the bucket
	log.Println("Fetching objects from the bucket...")
	for {
		attr, err := obj_it.Next()
		if err == iterator.Done {
			log.Println("All objects fetched successfully")
			break
		} else if err != nil {
			log.Printf("Error fetching object: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Found object: %s, Size: %d bytes, LastModified: %s", attr.Name, attr.Size, attr.Updated)
		objects = append(objects, cloudutils.BucketObject{
			Name:         attr.Name,
			Size:         attr.Size,
			LastModified: attr.Updated.String(),
		})
	}

	// Encode the objects to JSON
	w.Header().Set("Content-Type", "application/json")
	log.Println("Encoding objects to JSON...")

	if err := json.NewEncoder(w).Encode(objects); err != nil {
		log.Printf("Error encoding objects to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Objects successfully returned in JSON format")
}
