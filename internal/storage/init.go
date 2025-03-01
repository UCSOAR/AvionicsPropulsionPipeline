package storage

import "os"

var storageDirs = [...]string{
	DefaultUploadContext.BasePath,
	DefaultCacheContext.BasePath,
}

func InitStorageDirectories() error {
	for _, dir := range storageDirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
