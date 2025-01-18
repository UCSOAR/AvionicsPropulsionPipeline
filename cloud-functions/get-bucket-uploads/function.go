package function

import (
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/storage"
	cloudUtils "example.com/cloud-utils"
	"google.golang.org/api/iterator"
)

func GetBucketUploads(w http.ResponseWriter, r *http.Request) {
	cloudUtils.SetCorsHeaders(w, cloudUtils.Cors{
		AllowOrigin:  "*",
		AllowMethods: []string{"GET", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
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

	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		http.Error(w, "Failed to create GCS client", http.StatusInternalServerError)
		return
	}

	defer client.Close()

	bucket := client.Bucket(cloudUtils.BucketName)
	obj_it := bucket.Objects(ctx, nil)

	var objects []cloudUtils.BucketObject

	for {
		attr, err := obj_it.Next()

		if err == iterator.Done {
			break
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		objects = append(objects, cloudUtils.BucketObject{
			Name:         attr.Name,
			Size:         attr.Size,
			LastModified: attr.Updated.String(),
		})
	}

	// Encode the objects to JSON
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(objects); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
