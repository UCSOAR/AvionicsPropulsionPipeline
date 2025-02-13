package controllers

import (
	"bufio"
	"net/http"
	storage "soarpipeline/internal/storage"
	staticfire "soarpipeline/pkg/staticfire"
	"strings"
	"sync"
)

const maxFileSize = 10 << 26 // 671 MegaBytes
const extension = ".lvm"

func PostStaticFire(w http.ResponseWriter, r *http.Request) {
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
	if strings.HasSuffix(header.Filename, extension) {
		http.Error(w, "Invalid file extension", http.StatusBadRequest)
		return
	}

	// Create cache tree
	name := header.Filename[:len(header.Filename)-len(extension)]

	scanner := bufio.NewScanner(file)
	tree, err := staticfire.ParseIntoCacheTree(scanner)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store files
	errorChan := make(chan error, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		// Store cache tree
		if err = storage.DefaultCacheStorageContext.StoreTree(name, &tree); err != nil {
			errorChan <- err
		}
	}()

	go func() {
		defer wg.Done()

		// Store uploaded file
		if err = storage.DefaultUploadStorageContext.Store(header.Filename, file); err != nil {
			errorChan <- err
		}
	}()

	wg.Wait()
	close(errorChan)

	// Check for errors
	for err := range errorChan {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
