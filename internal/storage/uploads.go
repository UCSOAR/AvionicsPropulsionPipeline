package storage

import (
	"mime/multipart"
	"os"
	"path"
)

const uploadsSubdirName = "uploads"

// Store an uploaded file.
//
// Parameters:
//   - name: The name of the file to store.
//   - file: The file to store.
//
// Returns:
//   - error: An error if the file could not be stored, or nil if the operation was successful.
func StoreUpload(name string, file multipart.File) error {
	uploadPath := path.Join(StorageDirPath, uploadsSubdirName, name)
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

// Delete an uploaded file.
//
// Parameters:
//   - name: The name of the file to delete.
//
// Returns:
//   - error: An error if the file could not be deleted, or nil if the operation was successful.
func DeleteUpload(name string) error {
	uploadPath := path.Join(StorageDirPath, uploadsSubdirName, name)

	if err := os.Remove(uploadPath); err != nil {
		return err
	}

	return nil
}
