package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"soarpipeline/internal/storage"
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
	header := append(xCols, yCols...)
	csvWriter.Write(header)

	// Write rows
	rowCount := len(xColumns[xCols[0]].Rows)
	for i := 0; i < rowCount; i++ {
		row := []string{}
		for _, col := range xCols {
			row = append(row, fmt.Sprint(xColumns[col].Rows[i]))
		}
		for _, col := range yCols {
			row = append(row, fmt.Sprint(yColumns[col].Rows[i]))
		}
		csvWriter.Write(row)
	}
}
