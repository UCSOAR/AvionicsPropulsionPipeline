package storage

import (
	"fmt"
	"io"
	"os"
)

// Store a file.
//
// Parameters:
//   - name: The name of the file to store.
//   - reader: A reader for the contents of the file to store.
//
// Returns:
//   - error: An error if the file could not be stored, or nil if the operation was successful.
func (ctx *StorageContext) Store(name string, reader io.Reader) error {
	uploadPath := ctx.GetFilePath(name)

	// Check if the file already exists
	if _, err := os.Stat(uploadPath); err == nil {
		return os.ErrExist
	}

	dest, err := os.Create(uploadPath)

	if err != nil {
		return err
	}

	defer dest.Close()

	if len, err := io.Copy(dest, reader); err != nil {
		return err
	} else {
		fmt.Printf("Stored %s (%d bytes)\n", name, len)
	}

	return nil
}

// Delete a file.
//
// Parameters:
//   - name: The name of the file to delete.
//
// Returns:
//   - error: An error if the file could not be deleted, or nil if the operation was successful.
func (ctx *StorageContext) Delete(name string) error {
	uploadPath := ctx.GetFilePath(name)

	if err := os.Remove(uploadPath); err != nil {
		return err
	}

	return nil
}
