package storage

import (
	"mime/multipart"
	"os"
)

// Store a file.
//
// Parameters:
//   - name: The name of the file to store.
//   - file: The file to store.
//
// Returns:
//   - error: An error if the file could not be stored, or nil if the operation was successful.
func (ctx *StorageContext) Store(name string, file multipart.File) error {
	uploadPath := ctx.GetFilePath(name)
	uploadFile, err := os.Create(uploadPath)

	if err != nil {
		return err
	}

	defer uploadFile.Close()

	if _, err := uploadFile.ReadFrom(file); err != nil {
		return err
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
