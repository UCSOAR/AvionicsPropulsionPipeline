package function

import (
	"encoding/json"
	"net/http"

	staticFireCaching "github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/caching"
)

type GetStaticFireColumnsRequest struct {
	Name         string   `json:"name"`
	XColumnNames []string `json:"xColumnNames"`
	YColumnNames []string `json:"yColumnNames"`
}

type GetStaticFireColumnsResponse struct {
	XColumns []staticFireCaching.XColumnNode `json:"xColumns"`
	YColumns []staticFireCaching.YColumnNode `json:"yColumns"`
}

func GetStaticFireColumns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {
	case "GET":
		break // Only allow GET requests
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	default:
		http.Error(w, "Only GET requests allowed for fetching static fire columns", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req GetStaticFireColumnsRequest

	{
		reqDecoder := json.NewDecoder(r.Body)

		if err := reqDecoder.Decode(&req); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}
	}

	// Get columns
	xColumns, yColumns, err := staticFireCaching.GetColumns(req.Name, req.XColumnNames, req.YColumnNames)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write response
	res := GetStaticFireColumnsResponse{
		XColumns: xColumns,
		YColumns: yColumns,
	}

	resJson, err := json.Marshal(res)

	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}
