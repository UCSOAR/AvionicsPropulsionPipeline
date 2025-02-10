package storage

import (
	"encoding/gob"
	"os"
	"path"
	staticfire "soarpipeline/pkg/staticfire"
	"sync"
)

const cacheSubdirName = "cache"
const xColumnsSubdirName = "x"
const yColumnsSubdirName = "y"
const previewMetadataFileName = "preview"

// Helper function to store a gob object to a file.
//
// Parameters:
//   - path: The path to the file to store the gob object in.
//   - obj: The gob object to store.
//
// Returns:
//   - error: An error if the gob object could not be stored, or nil if the operation was successful.
func storeGobObject[T any](path string, obj *T) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()
	encoder := gob.NewEncoder(file)

	if err := encoder.Encode(obj); err != nil {
		return err
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
func decodeGobObject[T any](path string, emptyObj *T) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()
	decoder := gob.NewDecoder(file)

	if err := decoder.Decode(emptyObj); err != nil {
		return err
	}

	return nil
}

// Stores the cache tree data in the cache directory.
//
// Parameters:
//   - name: The name of the cache directory to store the cache tree data in.
//   - tree: The cache tree data to store.
//
// Returns:
//   - error: An error if the cache tree data could not be stored, or nil if the operation was successful.
func StoreCache(name string, tree *staticfire.CacheTree) error {
	cacheBasePath := path.Join(StorageDirPath, cacheSubdirName, name)
	xColumnsBasePath := path.Join(cacheBasePath, xColumnsSubdirName)
	yColumnsBasePath := path.Join(cacheBasePath, yColumnsSubdirName)

	// Create the cache directory and subdirectories
	if err := os.MkdirAll(xColumnsBasePath, os.ModePerm); err != nil {
		return err
	}

	if err := os.MkdirAll(yColumnsBasePath, os.ModePerm); err != nil {
		return err
	}

	// Start storing the cache tree data
	errorChan := make(chan error)
	wg := sync.WaitGroup{}

	// Store the preview metadata
	{
		previewMetadataPath := path.Join(cacheBasePath, previewMetadataFileName)
		wg.Add(1)

		go func() {
			defer wg.Done()

			if err := storeGobObject(previewMetadataPath, &tree.PreviewMetadata); err != nil {
				errorChan <- err
			}
		}()
	}

	// Store the X column data
	for i, xCol := range tree.XColumnNodes {
		xColPath := path.Join(xColumnsBasePath, tree.PreviewMetadata.XColumnNames[i])
		wg.Add(1)

		go func() {
			defer wg.Done()

			if err := storeGobObject(xColPath, &xCol); err != nil {
				errorChan <- err
			}
		}()
	}

	// Store the Y column data
	for j, yCol := range tree.YColumnNodes {
		yColPath := path.Join(yColumnsBasePath, tree.PreviewMetadata.YColumnNames[j])
		wg.Add(1)

		go func() {
			defer wg.Done()

			if err := storeGobObject(yColPath, &yCol); err != nil {
				errorChan <- err
			}
		}()
	}

	// Ensure error channel gets closed
	go func() {
		wg.Wait()

		close(errorChan)
	}()

	// Check for errors
	for errorChan != nil {
		select {
		case err, ok := <-errorChan:
			if !ok {
				errorChan = nil
			} else {
				return err
			}
		}
	}

	return nil
}

// Reads the preview metadata for all caches in the cache directory.
//
// Returns:
//   - metadata: A map of cache names to their respective preview metadata.
//   - error: An error if the metadata could not be read, or nil if the operation was successful.
func ReadAllCacheMetadata() (map[string]staticfire.PreviewMetadata, error) {
	type MetadataKV struct {
		name     string
		metadata staticfire.PreviewMetadata
	}

	// Look at all subdirectories in the cache directory
	cacheDirPath := path.Join(StorageDirPath, cacheSubdirName)
	entries, err := os.ReadDir(cacheDirPath)

	if err != nil {
		return nil, err
	}

	// Read the preview metadata for each cache
	metadata := make(map[string]staticfire.PreviewMetadata, len(entries))
	metadataChan := make(chan MetadataKV, len(entries))
	errorChan := make(chan error)
	wg := sync.WaitGroup{}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		cacheName := entry.Name()

		// Read the preview metadata
		previewMetadataPath := path.Join(cacheDirPath, cacheName, previewMetadataFileName)
		wg.Add(1)

		go func() {
			defer wg.Done()
			var metadata staticfire.PreviewMetadata

			if err := decodeGobObject(previewMetadataPath, &metadata); err != nil {
				errorChan <- err
			} else {
				metadataChan <- MetadataKV{cacheName, metadata}
			}
		}()
	}

	// Ensure channels are closed
	go func() {
		wg.Wait()

		close(metadataChan)
		close(errorChan)
	}()

	// Check for errors and collect metadata
	for metadataChan != nil || errorChan != nil {
		select {
		case kv, ok := <-metadataChan:
			if !ok {
				metadataChan = nil
			} else {
				metadata[kv.name] = kv.metadata
			}
		case err, ok := <-errorChan:
			if !ok {
				errorChan = nil
			} else {
				return nil, err
			}
		}
	}

	return metadata, nil
}

// Retrieves the X and Y columns for the specified cache.
// Each column is identified by its name in the cache tree.
//
// Parameters:
//   - name: The name of the cache directory to retrieve columns from.
//   - xColumnNames: The names of the X columns to retrieve.
//   - yColumnNames: The names of the Y columns to retrieve.
//
// Returns:
//   - xColumnNodes: A map of X column names to their respective column nodes.
//   - yColumnNodes: A map of Y column names to their respective column nodes.
//   - error: An error if the columns could not be retrieved, or nil if the operation was successful.
func GetCachedColumns(name string, xColumnNames []string, yColumnNames []string) (map[string]staticfire.XColumnNode, map[string]staticfire.YColumnNode, error) {
	type ColumnKV[T any] struct {
		name   string
		column T
	}

	xColumnNodes := make(map[string]staticfire.XColumnNode, len(xColumnNames))
	yColumnNodes := make(map[string]staticfire.YColumnNode, len(yColumnNames))

	// TODO: Implement this function
	_ = xColumnNodes
	_ = yColumnNodes

	return nil, nil, nil
}

// Removes the cache directory specified by the given name.
//
// Parameters:
//   - name: The name of the cache directory to be deleted.
//
// Returns:
//   - error: An error if the directory could not be removed, or nil if the operation was successful.
func DeleteCache(name string) error {
	cacheDirPath := path.Join(StorageDirPath, cacheSubdirName, name)

	if err := os.RemoveAll(cacheDirPath); err != nil {
		return err
	}

	return nil
}
