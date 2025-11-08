package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

func PostCreateFolder(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Type string `json:"type"`
		Path string `json:"path"`
		Name string `json:"name"`
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	baseDir := filepath.Join("storage/cache", req.Path)
	target := filepath.Join(baseDir, req.Name)

	switch req.Type {
	case "folder":
		err := os.MkdirAll(target, os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "file":
		f, err := os.Create(target)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
	default:
		http.Error(w, "Invalid type", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"folder": target,
	})

}
