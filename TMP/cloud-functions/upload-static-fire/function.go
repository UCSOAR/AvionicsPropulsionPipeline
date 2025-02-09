package function

import (
	"net/http"
	"strings"

	bucketInfo "github.com/UCSOAR/AvionicsPropulsionPipeline/bucket-info"
	staticFireParser "github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/parser"
)

func UploadStaticFire(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {
	case "POST":
		{
			// Only allow POST requests
			break
		}
	case "OPTIONS":
		{
			w.WriteHeader(http.StatusOK)
			return
		}
	default:
		{
			http.Error(w, "Only POST requests allowed for static fire uploads", http.StatusMethodNotAllowed)
			return
		}
	}

	// Limit size of uploaded file
	if r.ContentLength > bucketInfo.MaxFileSize {
		http.Error(w, "File is too large", http.StatusBadRequest)
		return
	}

	// Parse form data
	if err := r.ParseMultipartForm(bucketInfo.MaxFileSize); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get file from form data
	file, header, err := r.FormFile("file")
	fileName := header.Filename

	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Ensure file has `.lvm` extension
	if strings.HasSuffix(strings.ToLower(fileName), ".lvm") == false {
		http.Error(w, "File must have .lvm extension", http.StatusBadRequest)
		return
	}

	// Remove `.lvm` extension
	fileName = fileName[:len(fileName)-4]

	// Create the cache tree
	tree, err := staticFireParser.ParseIntoCacheTree(file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store the cache tree in GCS
	err = tree.Store(fileName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
