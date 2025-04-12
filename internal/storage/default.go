package storage

import "path"

// Path to where this project will store its data.
const StorageDirPath = "./storage"

// The storage context for LVM file uploads.
// SHOULD NOT BE MODIFIED AT RUNTIME.
//
//nolint:gochecknoglobals
var DefaultUploadContext = StorageContext{
	BasePath: path.Join(StorageDirPath, "uploads"),
}

// The storage context for cache trees generated from LVM files.
// SHOULD NOT BE MODIFIED AT RUNTIME.
//
//nolint:gochecknoglobals
var DefaultCacheContext = CacheStorageContext{
	StorageContext:             StorageContext{BasePath: path.Join(StorageDirPath, "cache")},
	YColumnsMetadataSubdirName: "y_metadata",
	XColumnsSubdirName:         "x_columns",
	YColumnsSubdirName:         "y_columns",
	PreviewMetadataFileName:    "preview",
}
