package controllers

import (
	"encoding/json"
	"net/http"
	storage "soarpipeline/internal/storage"
	staticfire "soarpipeline/pkg/staticfire"
)

// Represents the expected JSON structure of the request body
type PostStaticFireColumnsRequest struct {
	Name         string   `json:"name"`
	XColumnNames []string `json:"xColumnNames"`
	YColumnNames []string `json:"yColumnNames"`
}

// Represents the expected JSON structure of the response body
type PostStaticFireColumnsResponse struct {
	YColumnMetadata map[string]staticfire.YColumnMetadata `json:"yColumnMetadata"`
	XColumns        map[string]staticfire.ColumnNode      `json:"xColumns"`
	YColumns        map[string]staticfire.ColumnNode      `json:"yColumns"`
}

func PostStaticFireColumns(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req PostStaticFireColumnsRequest

	{
		reqDecoder := json.NewDecoder(r.Body)

		if err := reqDecoder.Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	// Get columns
	yColumnMetadata, xColumns, yColumns, err := storage.DefaultCacheContext.ReadColumns(req.Name, req.XColumnNames, req.YColumnNames)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create response
	res := PostStaticFireColumnsResponse{
		YColumnMetadata: yColumnMetadata,
		XColumns:        xColumns,
		YColumns:        yColumns,
	}

	// Write response
	resJson, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}
