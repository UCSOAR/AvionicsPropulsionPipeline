package storage

import "path"

// Path to where this project will store its data.
const storageDirPath = "./storage"

// The storage context for LVM file uploads.
var DefaultUploadStorageContext = StorageContext{
	BasePath: path.Join(storageDirPath, "uploads"),
}

// The storage context for cache trees generated from LVM files.
var DefaultCacheStorageContext = CacheStorageContext{
	StorageContext:          StorageContext{BasePath: path.Join(storageDirPath, "cache")},
	XColumnsSubdirName:      "x",
	YColumnsSubdirName:      "y",
	PreviewMetadataFileName: "preview",
}
