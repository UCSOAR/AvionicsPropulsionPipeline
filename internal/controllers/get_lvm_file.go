package controllers

import (
	"net/http"
	"soarpipeline/internal/storage"
)

var storeCtx = storage.DefaultUploadContext

func GetLVMFile(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	if filename == "" {
		http.Error(w, "missing file name", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
	w.Header().Set("Content-Type", "application/octet-stream")

	if err := storeCtx.DownloadLVM(filename, w); err != nil {
		http.Error(w, "could not download file", http.StatusInternalServerError)
		return
	}
}
