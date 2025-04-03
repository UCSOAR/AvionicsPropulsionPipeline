package controllers

import (
	"encoding/json"
	"net/http"
	"soarpipeline/internal/storage"
)

func GetStaticFireMetadata(w http.ResponseWriter, r *http.Request) {
	// Retrieve static fire metadata
	metadata, err := storage.DefaultCacheContext.ReadAllPreviewMetadata()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(metadata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
