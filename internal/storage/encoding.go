package storage

import (
	"encoding/binary"
	"encoding/gob"
	"io"
	"os"
	"soarpipeline/pkg/staticfire"
	"unsafe"
)

func EncodeColumnNode(path string, node *staticfire.ColumnNode) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	for _, row := range node.Rows {
		if err := binary.Write(file, binary.LittleEndian, row); err != nil {
			return err
		}
	}

	return nil
}

func DecodeColumnNode(path string, startRow int, numRows int) (staticfire.ColumnNode, error) {
	file, err := os.Open(path)

	if err != nil {
		return staticfire.ColumnNode{}, err
	}

	defer file.Close()

	const rowSize = int64(unsafe.Sizeof(float64(0)))

	{
		// Get the file size
		fileInfo, err := file.Stat()

		if err != nil {
			return staticfire.ColumnNode{}, err
		}

		fileSize := fileInfo.Size()

		// Calculate the total number of rows in the file
		totalRows := int(fileSize / rowSize)

		// Adjust number of rows to read if necessary
		if startRow+numRows > totalRows {
			numRows = totalRows - startRow
		}
	}

	// Seek to the position of the start row
	{
		startPos := int64(startRow) * rowSize

		if _, err := file.Seek(startPos, io.SeekStart); err != nil {
			return staticfire.ColumnNode{}, err
		}
	}

	// Read specified number of rows
	rows := make([]float64, numRows)

	if err := binary.Read(file, binary.LittleEndian, &rows); err != nil {
		return staticfire.ColumnNode{}, err
	}

	node := staticfire.ColumnNode{
		Rows: rows,
	}

	return node, nil
}

// Helper function to store a gob object to a file.
//
// Parameters:
//   - path: The path to the file to store the gob object in.
//   - obj: The gob object to store.
//
// Returns:
//   - error: An error if the gob object could not be stored, or nil if the operation was successful.
func EncodeGobObject[T any](path string, obj *T) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	{
		encoder := gob.NewEncoder(file)

		if err := encoder.Encode(obj); err != nil {
			return err
		}
	}

	return nil
}

// Helper function to decode a gob object from a file.
//
// Parameters:
//   - path: The path to the file containing the gob object.
//   - emptyObj: A pointer to an empty object of the type to decode.
//
// Returns:
//   - error: An error if the gob object could not be decoded, or nil if the operation was successful.
func DecodeGobObject[T any](path string, emptyObj *T) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	{
		decoder := gob.NewDecoder(file)

		if err := decoder.Decode(emptyObj); err != nil {
			return err
		}
	}

	return nil
}
