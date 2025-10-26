package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soarpipeline/pkg/filters"
)

// Represents the expected JSON structure of the request body.
type PostFilterDataRequest struct {
	XColumns     []float64 `json:"xColumns"`
	YColumns     []float64 `json:"yColumns"`
	FilterValue  float64   `json:"filterValue"`
	FilterNumber int       `json:"filterNumber"`
}

// Represents the expected JSON structure of the response body.
type PostFilterDataResponse struct {
	XColumns []float64 `json:"xColumns"`
	YColumns []float64 `json:"yColumns"`
}

func PostFilterData(w http.ResponseWriter, r *http.Request) {
	// Parse request body.
	var req PostFilterDataRequest

	{
		reqDecoder := json.NewDecoder(r.Body)

		if err := reqDecoder.Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	fmt.Println(req)

	if (req.FilterNumber) == 1 {

		yFilteredCol := filters.GaussianFilter(req.XColumns, req.YColumns, req.FilterValue)

		res := PostFilterDataResponse{
			XColumns: req.XColumns,
			YColumns: yFilteredCol,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	if (req.FilterNumber) == 2 {

		yFilteredCol := filters.MovingAvgFilter(req.XColumns, req.YColumns, int(req.FilterValue))

		res := PostFilterDataResponse{
			XColumns: req.XColumns,
			YColumns: yFilteredCol,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

}
