package storage

import (
	"os"

	"soarpipeline/pkg/staticfire"
	"sync"
)

// Stores the cache tree data in the cache directory.
//
// Parameters:
//   - name: The name of the cache directory to store the cache tree data in.
//   - tree: The cache tree data to store.
//
// Returns:
//   - error: An error if the cache tree data could not be stored, or nil if the operation was successful.
func (ctx *CacheStorageContext) StoreTree(name string, tree *staticfire.CacheTree) error {
	yColumnsMetadataPath := ctx.GetYColumnsMetadataPath(name)
	xColumnsPath := ctx.GetXColumnsPath(name)
	yColumnsPath := ctx.GetYColumnsPath(name)
	previewMetadataFile := ctx.GetPreviewMetadataFilePath(name)

	// Check if the cache directory already exists
	if _, err := os.Stat(ctx.GetCachePath(name)); err == nil {
		return os.ErrExist
	}

	// Create the cache directory and subdirectories
	{
		paths := [...]string{
			yColumnsMetadataPath,
			xColumnsPath,
			yColumnsPath,
		}

		for _, path := range paths {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
		}
	}

	// Start storing the cache tree data
	errorChan := make(chan error)
	wg := sync.WaitGroup{}

	// Store the preview metadata
	{
		wg.Add(1)

		go func() {
			defer wg.Done()

			if err := EncodeGobObject(previewMetadataFile, &tree.PreviewMetadata); err != nil {
				errorChan <- err
			}
		}()
	}

	// Store the Y column metadata
	for k, yColMeta := range tree.YColumnMetadata {
		yColMetaPath := ctx.GetYColumnMetadataFilePath(name, tree.PreviewMetadata.YColumnNames[k])
		wg.Add(1)

		go func() {
			defer wg.Done()

			if err := EncodeGobObject(yColMetaPath, &yColMeta); err != nil {
				errorChan <- err
			}
		}()
	}

	// Store the X column data
	for i, xCol := range tree.XColumnNodes {
		xColPath := ctx.GetXColumnFilePath(name, tree.PreviewMetadata.XColumnNames[i])
		wg.Add(1)

		go func() {
			defer wg.Done()

			if err := EncodeColumnNode(xColPath, &xCol); err != nil {
				errorChan <- err
			}
		}()
	}

	// Store the Y column data
	for j, yCol := range tree.YColumnNodes {
		yColPath := ctx.GetYColumnFilePath(name, tree.PreviewMetadata.YColumnNames[j])
		wg.Add(1)

		go func() {
			defer wg.Done()

			if err := EncodeColumnNode(yColPath, &yCol); err != nil {
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
		err, ok := <-errorChan

		if !ok {
			errorChan = nil
		} else {
			return err
		}
	}

	return nil
}

// Reads the preview metadata for all caches in the cache directory.
//
// Returns:
//   - metadata: A map of cache names to their respective preview metadata.
//   - error: An error if the metadata could not be read, or nil if the operation was successful.
func (ctx *CacheStorageContext) ReadAllPreviewMetadata() (map[string]staticfire.PreviewMetadata, error) {
	type MetadataKV struct {
		name     string
		metadata staticfire.PreviewMetadata
	}

	// Look at all subdirectories in the cache directory
	entries, err := os.ReadDir(ctx.BasePath)

	if err != nil {
		// Treat the cache directory as empty if it does not exist
		return make(map[string]staticfire.PreviewMetadata), nil
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
		previewMetadataPath := ctx.GetPreviewMetadataFilePath(cacheName)
		wg.Add(1)

		go func() {
			defer wg.Done()
			var metadata staticfire.PreviewMetadata

			if err := DecodeGobObject(previewMetadataPath, &metadata); err != nil {
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
//   - startRow: The index of the first row to retrieve.
//   - numRows: The number of rows to retrieve after the start row.
//   - xColumnNames: The names of the X columns to retrieve.
//   - yColumnNames: The names of the Y columns to retrieve.
//
// Returns:
//   - yColumnMetadata: A map of Y column names to their respective metadata.
//   - xColumnNodes: A map of X column names to their respective column nodes.
//   - yColumnNodes: A map of Y column names to their respective column nodes.
//   - error: An error if the columns could not be retrieved, or nil if the operation was successful.
func (ctx *CacheStorageContext) ReadColumns(name string, startRow int, numRows int, xColumnNames []string, yColumnNames []string) (map[string]staticfire.YColumnMetadata, map[string]staticfire.ColumnNode, map[string]staticfire.ColumnNode, error) {
	type Kv[T any] struct {
		key   string
		value T
	}

	// Retrieve columns and metadata concurrently
	errorChan := make(chan error)
	wg := sync.WaitGroup{}

	// Retrieve Y column metadata
	yColumnMetadataChan := make(chan Kv[staticfire.YColumnMetadata], len(yColumnNames))

	for _, yColName := range yColumnNames {
		yColMetaPath := ctx.GetYColumnMetadataFilePath(name, yColName)
		wg.Add(1)

		go func() {
			defer wg.Done()
			var yColMeta staticfire.YColumnMetadata

			if err := DecodeGobObject(yColMetaPath, &yColMeta); err != nil {
				errorChan <- err
			} else {
				yColumnMetadataChan <- Kv[staticfire.YColumnMetadata]{yColName, yColMeta}
			}
		}()
	}

	// Retrieve the X columns
	xColumnChan := make(chan Kv[staticfire.ColumnNode], len(xColumnNames))

	for _, xColName := range xColumnNames {
		xColPath := ctx.GetXColumnFilePath(name, xColName)
		wg.Add(1)

		go func() {
			defer wg.Done()

			if xCol, err := DecodeColumnNode(xColPath, startRow, numRows); err != nil {
				errorChan <- err
			} else {
				xColumnChan <- Kv[staticfire.ColumnNode]{xColName, xCol}
			}
		}()
	}

	// Retrieve the Y columns
	yColumnChan := make(chan Kv[staticfire.ColumnNode], len(yColumnNames))

	for _, yColName := range yColumnNames {
		yColPath := ctx.GetYColumnFilePath(name, yColName)
		wg.Add(1)

		go func() {
			defer wg.Done()

			if yCol, err := DecodeColumnNode(yColPath, startRow, numRows); err != nil {
				errorChan <- err
			} else {
				yColumnChan <- Kv[staticfire.ColumnNode]{yColName, yCol}
			}
		}()
	}

	// Ensure channels are closed
	go func() {
		wg.Wait()

		close(yColumnMetadataChan)
		close(xColumnChan)
		close(yColumnChan)
		close(errorChan)
	}()

	// Check for errors and collect columns
	yColumnMetadata := make(map[string]staticfire.YColumnMetadata, len(yColumnNames))
	xColumnNodes := make(map[string]staticfire.ColumnNode, len(xColumnNames))
	yColumnNodes := make(map[string]staticfire.ColumnNode, len(yColumnNames))

	for yColumnMetadataChan != nil || xColumnChan != nil || yColumnChan != nil || errorChan != nil {
		select {
		case yColMetaKv, ok := <-yColumnMetadataChan:
			if !ok {
				yColumnMetadataChan = nil
			} else {
				yColumnMetadata[yColMetaKv.key] = yColMetaKv.value
			}
		case xKv, ok := <-xColumnChan:
			if !ok {
				xColumnChan = nil
			} else {
				xColumnNodes[xKv.key] = xKv.value
			}
		case yKv, ok := <-yColumnChan:
			if !ok {
				yColumnChan = nil
			} else {
				yColumnNodes[yKv.key] = yKv.value
			}
		case err, ok := <-errorChan:
			if !ok {
				errorChan = nil
			} else {
				return nil, nil, nil, err
			}
		}
	}

	return yColumnMetadata, xColumnNodes, yColumnNodes, nil
}

// Removes the cache tree directory specified by the given name.
//
// Parameters:
//   - name: The name of the cache directory to be deleted.
//
// Returns:
//   - error: An error if the directory could not be removed, or nil if the operation was successful.
func (ctx *CacheStorageContext) DeleteTree(name string) error {
	cacheDirPath := ctx.GetCachePath(name)

	if err := os.RemoveAll(cacheDirPath); err != nil {
		return err
	}

	return nil
}
