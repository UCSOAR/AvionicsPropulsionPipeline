package controllers

import (
	"net/http"
	"soarpipeline/internal/storage"
	"soarpipeline/pkg/staticfire"
	"strings"
)

const maxFileSize = 10 << 26 // 671 MegaBytes
const extension = ".lvm"

func PostUploadStaticFire(w http.ResponseWriter, r *http.Request) {
	// Limit size of uploaded file
	if r.ContentLength > maxFileSize {
		http.Error(w, "Uploaded file is too large", http.StatusRequestEntityTooLarge)
		return
	}

	// Parse the form data
	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get file from form data
	file, header, err := r.FormFile("file")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Assert file extension
	if !strings.HasSuffix(header.Filename, extension) {
		http.Error(w, "Invalid file extension", http.StatusBadRequest)
		return
	}

	// Create cache tree
	name := header.Filename[:len(header.Filename)-len(extension)]
	tree, err := staticfire.ParseIntoCacheTree(file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store cache tree
	if err = storage.DefaultCacheContext.StoreTree(name, &tree); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Reset file reader
	if _, err = file.Seek(0, 0); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store uploaded file
	if err = storage.DefaultUploadContext.Store(header.Filename, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
