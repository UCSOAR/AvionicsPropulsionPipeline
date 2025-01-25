package function

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	functionUtils "example.com/function-utils"
)

// Creates a directory in GCS if it does not already exist
// Returns true if the directory already exists, false otherwise
func CreateDirectory(dirName string, ctx context.Context, client *storage.Client) (bool, error) {
	// Create an empty object with the directory name as the object nam
	obj := client.Bucket(functionUtils.BucketName).Object(dirName + "/")

	// Check if the directory already exists
	if _, err := obj.Attrs(ctx); err == nil {
		// Directory exists
		return true, nil
	} else if err != storage.ErrObjectNotExist {
		return false, fmt.Errorf("Failed to check directory existence: %v", err)
	}

	// Directory does not exist, proceed to create
	writer := obj.NewWriter(ctx)

	if err := writer.Close(); err != nil {
		return false, fmt.Errorf("Failed to close directory GCS writer: %v", err)
	}

	return false, nil
}

func CreateObjectInDir(objectName string, path string, data []byte, ctx context.Context, client *storage.Client) error {
	// Assert that the path ends with a forward slash
	if path[:len(path)-1] != "/" {
		panic("Path must end with a forward slash")
	}

	bucket := client.Bucket(functionUtils.BucketName)
	obj := bucket.Object(path + objectName)

	// Check if the file exists
	if _, err := obj.Attrs(ctx); err == nil {
		// File exists
		return fmt.Errorf("File already exists")
	} else if err != storage.ErrObjectNotExist {
		return fmt.Errorf("Failed to check file existence: %v", err)
	}

	// File does not exist, proceed to upload
	writer := obj.NewWriter(ctx)

	// Copy file to GCS
	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("Failed to write data to GCS: %v", err)
	}

	// Close writer
	if err := writer.Close(); err != nil {
		return fmt.Errorf("Failed to close GCS writer: %v", err)
	}

	return nil
}
