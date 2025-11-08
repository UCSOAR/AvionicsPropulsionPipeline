package storage

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"soarpipeline/pkg/staticfire"
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
//
// Reads the preview metadata for all caches in the cache directory, recursively.
// This version supports subfolders (e.g. /cache/TestRuns/2025/RHT_2025-02-12_TestRun).
func (ctx *CacheStorageContext) ReadAllPreviewMetadata() (map[string]*staticfire.PreviewMetadata, error) {
	type MetadataKV struct {
		name     string
		metadata *staticfire.PreviewMetadata
	}

	basePath := ctx.BasePath
	metadata := make(map[string]*staticfire.PreviewMetadata)
	metadataChan := make(chan MetadataKV)
	errorChan := make(chan error)
	wg := sync.WaitGroup{}

	// Start directory walk
	go func() {
		filepath.WalkDir(basePath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				errorChan <- err
				return err
			}
			if !d.IsDir() {
				return nil
			}

			// Get relative path (skip base itself)
			rel, err := filepath.Rel(basePath, path)
			if err != nil || rel == "." {
				return nil
			}

			// Always add folder with nil metadata first
			folderKey := rel + string(os.PathSeparator)
			metadataChan <- MetadataKV{name: folderKey, metadata: nil}

			// Check if this directory has a preview file
			previewPath := filepath.Join(path, "preview")
			if _, err := os.Stat(previewPath); err == nil {
				wg.Add(1)
				go func(previewPath string, folderKey string) {
					defer wg.Done()
					var m staticfire.PreviewMetadata
					if err := DecodeGobObject(previewPath, &m); err != nil {
						errorChan <- fmt.Errorf("decode preview at %s: %w", previewPath, err)
						return
					}
					// Treat the preview as the "file" under this folder
					fileKey := strings.TrimSuffix(folderKey, string(os.PathSeparator))
					metadataChan <- MetadataKV{name: fileKey, metadata: &m}
				}(previewPath, folderKey)

				// Skip internal folders (x_columns, y_columns, etc.)
				return filepath.SkipDir
			}
			return nil
		})

		wg.Wait()
		close(metadataChan)
		close(errorChan)
	}()

	// Collect results
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

// GetTree reads the cached preview + all X/Y columns and rebuilds a CacheTree.
func (ctx *CacheStorageContext) GetTree(name string) (*staticfire.CacheTree, error) {
	// 1) Read preview metadata (you stored it as a gob in "preview")
	var preview staticfire.PreviewMetadata
	if err := DecodeGobObject(ctx.GetPreviewMetadataFilePath(name), &preview); err != nil {
		return nil, fmt.Errorf("read preview metadata: %w", err)
	}

	tree := &staticfire.CacheTree{
		PreviewMetadata: preview,
		// NOTE: it's fine to leave YColumnMetadata zeroed if you don't need it for Excel,
		// but we can load it too (next block).
		XColumnNodes:    make([]staticfire.ColumnNode, len(preview.XColumnNames)),
		YColumnNodes:    make([]staticfire.ColumnNode, len(preview.YColumnNames)),
		YColumnMetadata: make([]staticfire.YColumnMetadata, len(preview.YColumnNames)),
	}

	// 2) Read Y column metadata (one gob per column name)
	for i, yName := range preview.YColumnNames {
		var ymeta staticfire.YColumnMetadata
		if err := DecodeGobObject(ctx.GetYColumnMetadataFilePath(name, yName), &ymeta); err != nil {
			return nil, fmt.Errorf("read y metadata %q: %w", yName, err)
		}
		tree.YColumnMetadata[i] = ymeta
	}

	// helper: compute total rows from file size (float64 = 8 bytes) and decode all rows
	readAll := func(filePath string) (staticfire.ColumnNode, error) {
		fi, err := os.Stat(filePath)
		if err != nil {
			return staticfire.ColumnNode{}, fmt.Errorf("stat %s: %w", filePath, err)
		}
		totalRows := int(fi.Size() / 8) // float64 is 8 bytes, your encoder writes raw float64s
		if totalRows < 0 {
			totalRows = 0
		}
		node, err := DecodeColumnNode(filePath, 0, totalRows)
		if err != nil {
			return staticfire.ColumnNode{}, fmt.Errorf("decode %s: %w", filePath, err)
		}
		return node, nil
	}

	// 3) Load X columns
	for i, xName := range preview.XColumnNames {
		p := ctx.GetXColumnFilePath(name, xName) // e.g. ./storage/cache/<run>/x_columns/<name>
		node, err := readAll(p)
		if err != nil {
			return nil, err
		}
		tree.XColumnNodes[i] = node
	}

	// 4) Load Y columns
	for i, yName := range preview.YColumnNames {
		p := ctx.GetYColumnFilePath(name, yName) // e.g. ./storage/cache/<run>/y_columns/<name>
		node, err := readAll(p)
		if err != nil {
			return nil, err
		}
		tree.YColumnNodes[i] = node
	}

	return tree, nil
}
