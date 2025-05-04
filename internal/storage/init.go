package storage

import "os"

func InitStorageDirectories() error {
	storageDirs := [...]string{
		DefaultUploadContext.BasePath,
		DefaultCacheContext.BasePath,
	}

	for _, dir := range storageDirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
