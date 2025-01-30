package function

import (
	"context"
	"net/http"
	"strings"

	"cloud.google.com/go/storage"
	cloudFunctions "github.com/UCSOAR/AvionicsPropulsionPipeline/cloud-functions"
	staticFireParser "github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/parser"
)

func StaticFireUpload(w http.ResponseWriter, r *http.Request) {
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
	if r.ContentLength > cloudFunctions.MaxFileSize {
		http.Error(w, "File is too large", http.StatusBadRequest)
		return
	}

	// Parse form data
	if err := r.ParseMultipartForm(cloudFunctions.MaxFileSize); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get file from form data
	file, header, err := r.FormFile("file")
	objectName := header.Filename

	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Ensure file has `.lvm` extension
	if strings.HasSuffix(strings.ToLower(objectName), ".lvm") == false {
		http.Error(w, "File must have .lvm extension", http.StatusBadRequest)
		return
	}

	// Create a GCS client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		http.Error(w, "Failed to create GCS client", http.StatusInternalServerError)
		return
	}

	defer client.Close()

	// Assume we have parsed LVM for now
	lvm := staticFireParser.ParsedLvm{} // For now...
	cacheTree, err := lvm.ToCacheTree() // Right now, this returns a test cache tree

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Create the cache tree file structure
	_ = cacheTree // For now...

	// Create processed directory if it does not exist
	// if _, err := CreateDirectory(staticfire.ProcessedUploadsObjectName, ctx, client); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// // Create a directory to store processed file data
	// newDirName := objectName[:len(objectName)-4]
	// newDirExists, err := CreateDirectory(staticfire.ProcessedUploadsObjectName+newDirName, ctx, client)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// if newDirExists {
	// 	http.Error(w, "A file with this name has already been processed", http.StatusBadRequest)
	// 	return
	// }

	// // Read file data
	// rawFileText, err := io.ReadAll(file)

	// if err != nil {
	// 	http.Error(w, "Failed to read file data", http.StatusInternalServerError)
	// 	return
	// }

	// lvm, err := staticfireParser.ParseMainLvm(string(rawFileText))

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// // Create metadata for the new file
	// lvmMetadata := staticFireMetadata.LvmMetadata{
	// 	ProcessedTimestamp: time.Now().Format(time.RFC3339),
	// 	Operator:           lvm.EntryHeader.Operator,
	// 	ColumnNames:        lvm.SvData.ColumnNames,
	// }

	// // Encode parsed LVM and metadata into a binary format
	// lvmBytes, err := functionUtilsEncoding.BinaryEncode(lvm)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// lvmMetadataBytes, err := functionUtilsEncoding.BinaryEncode(lvmMetadata)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// // Write the LVM and metadata to GCS objects
	// if err := CreateObjectInDir(objectName, staticfire.ProcessedUploadsObjectName+newDirName+"/", lvmBytes, ctx, client); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// if err := CreateObjectInDir("metadata", staticfire.ProcessedUploadsObjectName+newDirName+"/", lvmMetadataBytes, ctx, client); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
}
