package controllers

import (
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"

	"soarpipeline/internal/storage"
	"soarpipeline/pkg/staticfire"
)

// ---- helpers ---------------------------------------------------------------

func openFloatColumn(path string, startRow int) (*os.File, int, error) {
	const float64Size = 8
	f, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	info, err := f.Stat()
	if err != nil {
		_ = f.Close()
		return nil, 0, err
	}
	totalRows := int(info.Size() / float64Size)
	if startRow > 0 {
		if _, err := f.Seek(int64(startRow*float64Size), io.SeekStart); err != nil {
			_ = f.Close()
			return nil, 0, err
		}
	}
	return f, totalRows, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ---- handler ---------------------------------------------------------------

// GET /api/staticfire/downloadExcel?file=<name>.lvm
// Optional: ?y=Y1,Y2   (limit Y columns, keeps preview order)
func GetExcelFromCache(w http.ResponseWriter, r *http.Request) {
	fileParam := r.URL.Query().Get("file")
	if fileParam == "" {
		http.Error(w, "missing file parameter", http.StatusBadRequest)
		return
	}
	name := strings.TrimSuffix(fileParam, filepath.Ext(fileParam))

	// Read preview metadata (column names)
	var pv staticfire.PreviewMetadata
	if err := storage.DecodeGobObject(storage.DefaultCacheContext.GetPreviewMetadataFilePath(name), &pv); err != nil {
		http.Error(w, fmt.Sprintf("failed to read preview metadata: %v", err), http.StatusNotFound)
		return
	}
	if len(pv.XColumnNames) == 0 {
		http.Error(w, "no X columns in cache", http.StatusBadRequest)
		return
	}
	xName := pv.XColumnNames[0]

	// Y selection
	var yNames []string
	if q := r.URL.Query().Get("y"); q != "" {
		want := map[string]struct{}{}
		for _, s := range strings.Split(q, ",") {
			if s = strings.TrimSpace(s); s != "" {
				want[s] = struct{}{}
			}
		}
		for _, yn := range pv.YColumnNames {
			if _, ok := want[yn]; ok {
				yNames = append(yNames, yn)
			}
		}
	} else {
		yNames = append(yNames, pv.YColumnNames...)
	}
	if len(yNames) == 0 {
		http.Error(w, "no Y columns selected", http.StatusBadRequest)
		return
	}

	// Open column files for streaming
	xPath := storage.DefaultCacheContext.GetXColumnFilePath(name, xName)
	xFile, totalRows, err := openFloatColumn(xPath, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("open X column: %v", err), http.StatusInternalServerError)
		return
	}
	defer xFile.Close()

	yFiles := make([]*os.File, len(yNames))
	for i, yn := range yNames {
		p := storage.DefaultCacheContext.GetYColumnFilePath(name, yn)
		f, _, err := openFloatColumn(p, 0)
		if err != nil {
			http.Error(w, fmt.Sprintf("open Y column %q: %v", yn, err), http.StatusInternalServerError)
			return
		}
		yFiles[i] = f
		defer f.Close()
	}

	// Build XLSX with StreamWriter and sheet rollover
	const (
		firstSheetName  = "Data"
		maxRowsPerSheet = 1_048_576 // Excel limit
		headerRowIndex  = 1
		firstDataRow    = headerRowIndex + 1
		chunk           = 8192 // rows per read/write batch
	)
	f := excelize.NewFile()

	// Sheet management
	currentSheetName := firstSheetName
	if _, err := f.NewSheet(currentSheetName); err != nil {
		http.Error(w, fmt.Sprintf("new sheet: %v", err), http.StatusInternalServerError)
		return
	}
	_ = f.DeleteSheet("Sheet1") // ignore if absent

	makeHeader := func(sw *excelize.StreamWriter) error {
		header := make([]interface{}, 0, 1+len(yNames))
		header = append(header, xName)
		for _, yn := range yNames {
			header = append(header, yn)
		}
		return sw.SetRow("A1", header)
	}

	sw, err := f.NewStreamWriter(currentSheetName)
	if err != nil {
		http.Error(w, fmt.Sprintf("stream writer: %v", err), http.StatusInternalServerError)
		return
	}
	if err := makeHeader(sw); err != nil {
		http.Error(w, fmt.Sprintf("write header: %v", err), http.StatusInternalServerError)
		return
	}

	currentRowInSheet := firstDataRow
	sheetIndex := 1

	rollover := func() error {
		if err := sw.Flush(); err != nil {
			return fmt.Errorf("flush stream: %w", err)
		}
		sheetIndex++
		currentSheetName = fmt.Sprintf("Data_%d", sheetIndex)
		if _, err := f.NewSheet(currentSheetName); err != nil {
			return fmt.Errorf("new sheet: %w", err)
		}
		var err error
		sw, err = f.NewStreamWriter(currentSheetName)
		if err != nil {
			return fmt.Errorf("stream writer: %w", err)
		}
		if err := makeHeader(sw); err != nil {
			return fmt.Errorf("write header: %w", err)
		}
		currentRowInSheet = firstDataRow
		return nil
	}

	// Buffers
	xBuf := make([]float64, chunk)
	yBufs := make([][]float64, len(yFiles))
	for i := range yBufs {
		yBufs[i] = make([]float64, chunk)
	}

	left := totalRows
	for left > 0 {
		select {
		case <-r.Context().Done():
			return // client aborted
		default:
		}

		n := min(chunk, left)

		// read X chunk
		if err := binary.Read(xFile, binary.LittleEndian, xBuf[:n]); err != nil {
			http.Error(w, fmt.Sprintf("read X: %v", err), http.StatusInternalServerError)
			return
		}
		// read Y chunks
		for i, yf := range yFiles {
			if err := binary.Read(yf, binary.LittleEndian, yBufs[i][:n]); err != nil {
				http.Error(w, fmt.Sprintf("read Y[%d]: %v", i, err), http.StatusInternalServerError)
				return
			}
		}

		for i := 0; i < n; i++ {
			// roll to new sheet if we’d exceed Excel’s limit
			if currentRowInSheet > maxRowsPerSheet {
				if err := rollover(); err != nil {
					http.Error(w, fmt.Sprintf("rollover: %v", err), http.StatusInternalServerError)
					return
				}
			}

			row := make([]interface{}, 0, 1+len(yBufs))
			row = append(row, xBuf[i])
			for j := range yBufs {
				row = append(row, yBufs[j][i])
			}
			cell, _ := excelize.CoordinatesToCellName(1, currentRowInSheet)
			if err := sw.SetRow(cell, row); err != nil {
				http.Error(w, fmt.Sprintf("stream row %d: %v", currentRowInSheet, err), http.StatusInternalServerError)
				return
			}
			currentRowInSheet++
		}

		left -= n
	}

	if err := sw.Flush(); err != nil {
		http.Error(w, fmt.Sprintf("flush stream: %v", err), http.StatusInternalServerError)
		return
	}

	// Send response
	filename := name + ".xlsx"
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Connection", "keep-alive")

	if err := f.Write(w); err != nil {
		http.Error(w, fmt.Sprintf("write xlsx: %v", err), http.StatusInternalServerError)
		return
	}
}
