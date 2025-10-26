package storage

import (
	"io"
	"os"
)

// Download a file from storage and write it to the given writer.
//
// Parameters:
//   - name: The name of the file to download.
//   - writer: The destination writer (e.g., http.ResponseWriter).
//
// Returns:
//   - error: If the file does not exist or cannot be read.
func (ctx *StorageContext) DownloadLVM(name string, writer io.Writer) error {
	filePath := ctx.GetFilePath(name)

	// open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// stream the file content to the writer
	_, err = io.Copy(writer, file)
	return err
}
