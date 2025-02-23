package storage

import "path"

// Represents the context for a storage operation.
// Contains a path to the base directory for the storage operation.
type StorageContext struct {
	BasePath string
}

// Get the full path to a file in the storage context.
//
// Parameters:
//   - name: The name of the file.
//
// Returns:
//   - string: The full path to the file.
func (ctx *StorageContext) GetFilePath(name string) string {
	return path.Join(ctx.BasePath, name)
}

// Represents the context for a cache storage operation.
// Contains metadata for the cache storage operation.
type CacheStorageContext struct {
	StorageContext
	YColumnsMetadataSubdirName string
	XColumnsSubdirName         string
	YColumnsSubdirName         string
	PreviewMetadataFileName    string
}

// Get the full path to a cache tree directory in the cache storage context.
//
// Parameters:
//   - name: The name of the cache tree directory.
//
// Returns:
//   - string: The full path to the cache tree directory.
func (ctx *CacheStorageContext) GetCachePath(name string) string {
	return path.Join(ctx.BasePath, name)
}

// Get the full path to the Y columns metadata directory in the cache storage context.
//
// Parameters:
//   - name: The name of the cache tree directory.
//
// Returns:
//   - string: The full path to the Y columns metadata directory.
func (ctx *CacheStorageContext) GetYColumnsMetadataPath(name string) string {
	return path.Join(ctx.BasePath, name, ctx.YColumnsMetadataSubdirName)
}

// Get the full path to the X columns directory in the cache storage context.
//
// Parameters:
//   - name: The name of the cache tree directory.
//
// Returns:
//   - string: The full path to the X columns directory.
func (ctx *CacheStorageContext) GetXColumnsPath(name string) string {
	return path.Join(ctx.BasePath, name, ctx.XColumnsSubdirName)
}

// Get the full path to the Y columns directory in the cache storage context.
//
// Parameters:
//   - name: The name of the cache tree directory.
//
// Returns:
//   - string: The full path to the Y columns directory.
func (ctx *CacheStorageContext) GetYColumnsPath(name string) string {
	return path.Join(ctx.BasePath, name, ctx.YColumnsSubdirName)
}

// Get the full path to the preview metadata file in the cache storage context.
//
// Parameters:
//   - name: The name of the cache tree directory.
//
// Returns:
//   - string: The full path to the preview metadata file.
func (ctx *CacheStorageContext) GetPreviewMetadataFilePath(name string) string {
	return path.Join(ctx.BasePath, name, ctx.PreviewMetadataFileName)
}

// Get the full path to a Y column metadata file in the cache storage context.
//
// Parameters:
//   - name: The name of the cache tree directory.
//   - yColName: The name of the Y column metadata file.
//
// Returns:
//   - string: The full path to the Y column metadata file.
func (ctx *CacheStorageContext) GetYColumnMetadataFilePath(name string, yColName string) string {
	return path.Join(ctx.BasePath, name, ctx.YColumnsMetadataSubdirName, yColName)
}

// Get the full path to an X column file in the cache storage context.
//
// Parameters:
//   - name: The name of the cache tree directory.
//   - xColName: The name of the X column file.
//
// Returns:
//   - string: The full path to the X column file.
func (ctx *CacheStorageContext) GetXColumnFilePath(name string, xColName string) string {
	return path.Join(ctx.BasePath, name, ctx.XColumnsSubdirName, xColName)
}

// Get the full path to a Y column file in the cache storage context.
//
// Parameters:
//   - name: The name of the cache tree directory.
//   - yColName: The name of the Y column file.
//
// Returns:
//   - string: The full path to the Y column file.
func (ctx *CacheStorageContext) GetYColumnFilePath(name string, yColName string) string {
	return path.Join(ctx.BasePath, name, ctx.YColumnsSubdirName, yColName)
}
