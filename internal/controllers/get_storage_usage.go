package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"soarpipeline/internal/storage"
)

// StorageSizeResponse represents the expected JSON structure of the response body
type StorageSizeResponse struct {
	StorageSizeBytes int64 `json:"storage_size_bytes"`
}

// CalculateStorageUsage calculates the total size of the storage directory
func CalculateStorageUsage(path string) (int64, error) {
	var size int64

	// Walk through the directory and sum up file sizes
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		size += info.Size()
		return nil
	})

	return size, err
}

// GetStorageUsage handles HTTP requests to retrieve storage size
func GetStorageUsage(w http.ResponseWriter, r *http.Request) {
	// Get the storage path
	storagePath := storage.StorageDirPath

	// Calculate the storage usage
	size, err := CalculateStorageUsage(storagePath)
	if err != nil {
		http.Error(w, "failed to calculate storage size", http.StatusInternalServerError)
		return
	}

	// Create response object
	response := StorageSizeResponse{StorageSizeBytes: size}

	// Set response header and encode response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
