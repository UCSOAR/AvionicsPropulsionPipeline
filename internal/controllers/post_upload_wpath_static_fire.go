package controllers

import (
	"net/http"
	"path/filepath"
	"soarpipeline/internal/storage"
	"soarpipeline/pkg/staticfire"
	"strings"
)

const maxFileSize = 10 << 26 // 671 MB
const extension = ".lvm"

// PostUploadPathStaticFire
// Uploads an .lvm file into a specified folder path from the frontend.
// The "path" form field can be empty (for root) or nested (e.g. "TestFolder/YawnTiredHelp").
func PostUploadPathStaticFire(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength > maxFileSize {
		http.Error(w, "Uploaded file is too large", http.StatusRequestEntityTooLarge)
		return
	}

	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Destination folder from UI — can be empty or nested (root if empty)
	destPath := strings.TrimSpace(r.FormValue("path"))

	// Extract the uploaded file
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file type
	if !strings.HasSuffix(strings.ToLower(header.Filename), extension) {
		http.Error(w, "Invalid file extension: only .lvm files are allowed", http.StatusBadRequest)
		return
	}

	baseName := header.Filename[:len(header.Filename)-len(extension)]

	// Parse into cache tree
	tree, err := staticfire.ParseIntoCacheTree(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Rewind file before storing
	if _, err = file.Seek(0, 0); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Build relative cache and upload paths using selected folder
	cacheRel := filepath.Join(destPath, baseName)
	if err = storage.DefaultCacheContext.StoreTree(cacheRel, &tree); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploadRel := filepath.Join(destPath, header.Filename)
	if err = storage.DefaultUploadContext.Store(uploadRel, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("File uploaded successfully to " + uploadRel))
}
