package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"soarpipeline/internal/storage"
)


// calculates total size of storage directory
func getStorageSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		size += info.Size()
		return nil
	})
	return size, err
}

// GetStorageSizeHandler handles HTTP requests to get storage size
func GetStorageSizeHandler(w http.ResponseWriter, r *http.Request) {
	storagePath := storage.StorageDirPath
	size, err := getStorageSize(storagePath)
	if err != nil {
		http.Error(w, "Failed to calculate storage size", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"storage_size_bytes": %d}`, size)
}
