package storage

import "path"

// Path to where this project will store its data.
const storageDirPath = "./storage"

// The storage context for LVM file uploads.
// SHOULD NOT BE MODIFIED AT RUNTIME.
var DefaultUploadContext = StorageContext{
	BasePath: path.Join(storageDirPath, "uploads"),
}

// The storage context for cache trees generated from LVM files.
// SHOULD NOT BE MODIFIED AT RUNTIME.
var DefaultCacheContext = CacheStorageContext{
	StorageContext:             StorageContext{BasePath: path.Join(storageDirPath, "cache")},
	YColumnsMetadataSubdirName: "y_metadata",
	XColumnsSubdirName:         "x_columns",
	YColumnsSubdirName:         "y_columns",
	PreviewMetadataFileName:    "preview",
}
