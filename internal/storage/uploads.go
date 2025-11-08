package storage

import (
	"io"
	"os"
	"path/filepath"
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
	// ✅ Normalize incoming path to OS format
	safeName := filepath.FromSlash(name)
	uploadPath := ctx.GetFilePath(safeName)

	if err := os.MkdirAll(filepath.Dir(uploadPath), 0755); err != nil {
		return err
	}

	if _, err := os.Stat(uploadPath); err == nil {
		return os.ErrExist
	}

	dest, err := os.Create(uploadPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	if _, err := io.Copy(dest, reader); err != nil {
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

func (ctx *StorageContext) StoreWithPath(folder string, name string, reader io.Reader) error {
	// ✅ Normalize incoming path to OS format
	safeFolder := filepath.FromSlash(folder)
	uploadDir := ctx.GetFilePath(safeFolder)

	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return err
	}

	uploadPath := filepath.Join(uploadDir, name)

	if _, err := os.Stat(uploadPath); err == nil {
		return os.ErrExist
	}

	dest, err := os.Create(uploadPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	if _, err := io.Copy(dest, reader); err != nil {
		return err
	}

	return nil
}
