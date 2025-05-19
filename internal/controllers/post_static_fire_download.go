package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"soarpipeline/internal/storage"
	"strconv"
)

func PostStaticFireDownload(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	startRow, _ := strconv.Atoi(r.URL.Query().Get("startRow"))
	numRows, _ := strconv.Atoi(r.URL.Query().Get("numRows"))
	xCols := r.URL.Query()["xColumnNames"] // multiple xColumnNames
	yCols := r.URL.Query()["yColumnNames"] // multiple yColumnNames

	_, xColumns, yColumns, err := storage.DefaultCacheContext.ReadColumns(name, startRow, numRows, xCols, yCols)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare CSV
	w.Header().Set("Content-Disposition", "attachment; filename=filtered_data.csv")
	w.Header().Set("Content-Type", "text/csv")

	csvWriter := csv.NewWriter(w)
	defer csvWriter.Flush()

	// Write header
	header := make([]string, 0, len(xCols)+len(yCols))
	header = append(header, xCols...)
	header = append(header, yCols...)

	if err := csvWriter.Write(header); err != nil {
		http.Error(w, "failed to write CSV header", http.StatusInternalServerError)
		return
	}

	// Write rows
	for i := range xColumns[xCols[0]].Rows {
		row := make([]string, 0, len(header))
		for _, col := range xCols {
			row = append(row, fmt.Sprint(xColumns[col].Rows[i]))
		}
		for _, col := range yCols {
			row = append(row, fmt.Sprint(yColumns[col].Rows[i]))
		}
		if err := csvWriter.Write(row); err != nil {
			http.Error(w, "failed to write CSV row", http.StatusInternalServerError)
			return	
		}
	}
}
