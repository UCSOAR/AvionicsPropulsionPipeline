package cachetree

// import (
// 	"context"
// 	"encoding/gob"
// 	"fmt"

// 	"cloud.google.com/go/storage"
// 	bucketInfo "github.com/UCSOAR/AvionicsPropulsionPipeline/bucket-info"
// 	"google.golang.org/api/iterator"
// )

// const xColumnsSubdir = "x"
// const yColumnsSubdir = "y"
// const previewMetadataFile = "preview"

// func createDirectory(ctx context.Context, bucket *storage.BucketHandle, name string) error {
// 	dirObj := bucket.Object(name + "/")

// 	if _, err := dirObj.Attrs(ctx); err == nil {
// 		// Directory already exists
// 		return fmt.Errorf("Directory with name %s already exists", name)
// 	} else if err != storage.ErrObjectNotExist {
// 		return fmt.Errorf("Failed to check directory existence: %v", err)
// 	}

// 	// Create the directory
// 	if err := dirObj.NewWriter(ctx).Close(); err != nil {
// 		return fmt.Errorf("Failed to create directory: %v", err)
// 	}

// 	return nil
// }

// func writeEncodedObject[T any](ctx context.Context, bucket *storage.BucketHandle, name string, obj T) error {
// 	writer := bucket.Object(name).NewWriter(ctx)
// 	encoder := gob.NewEncoder(writer)

// 	if err := encoder.Encode(obj); err != nil {
// 		return fmt.Errorf("Failed to encode object: %v", err)
// 	}

// 	if err := writer.Close(); err != nil {
// 		return fmt.Errorf("Failed to write object: %v", err)
// 	}

// 	return nil
// }

// func (tree *CacheTree) Store(name string) error {
// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx)

// 	if err != nil {
// 		return fmt.Errorf("Failed to create GCS client: %v", err)
// 	}

// 	defer client.Close()

// 	bucket := client.Bucket(bucketInfo.BucketName)

// 	// Create the tree structure in the file system
// 	// Create the root directory
// 	if err := createDirectory(ctx, bucket, name); err != nil {
// 		return err
// 	}

// 	// Create subdirectories for storing X and Y column data
// 	for _, dir := range []string{xColumnsSubdir, yColumnsSubdir} {
// 		if err := createDirectory(ctx, bucket, name+"/"+dir); err != nil {
// 			return err
// 		}
// 	}

// 	// Encode preview metadata and store it in the root directory
// 	if err := writeEncodedObject(ctx, bucket, name+"/"+previewMetadataFile, tree.PreviewMetadata); err != nil {
// 		return err
// 	}

// 	// Encode all X column data and store it in the X column directory
// 	for i, xCol := range tree.XColumnNodes {
// 		if err := writeEncodedObject(ctx, bucket, name+"/"+xColumnsSubdir+"/"+tree.PreviewMetadata.XColumnNames[i], xCol); err != nil {
// 			return err
// 		}
// 	}

// 	// Encode all Y column data and store it in the Y column directory
// 	for j, yCol := range tree.YColumnNodes {
// 		if err := writeEncodedObject(ctx, bucket, name+"/"+yColumnsSubdir+"/"+tree.PreviewMetadata.YColumnNames[j], yCol); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func GetAllCacheMetadata() (map[string]PreviewMetadata, error) {
// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx)

// 	if err != nil {
// 		return nil, fmt.Errorf("Failed to create GCS client: %v", err)
// 	}

// 	defer client.Close()

// 	bucket := client.Bucket(bucketInfo.BucketName)
// 	query := storage.Query{
// 		Delimiter:                "/",
// 		IncludeTrailingDelimiter: true,
// 	}

// 	query.SetAttrSelection([]string{"Prefix"})
// 	objIt := bucket.Objects(ctx, &query)

// 	// Map of cache names to preview metadata
// 	data := make(map[string]PreviewMetadata)

// 	for {
// 		attr, err := objIt.Next()

// 		if err == iterator.Done {
// 			break
// 		} else if err != nil {
// 			return nil, fmt.Errorf("Failed to iterate over objects: %v", err)
// 		}

// 		if attr.Prefix == "" {
// 			// Skip objects that are not directories
// 			continue
// 		}

// 		// Decode the preview metadata
// 		obj := bucket.Object(attr.Prefix + previewMetadataFile)
// 		reader, err := obj.NewReader(ctx)

// 		if err != nil {
// 			return nil, fmt.Errorf("Failed to read preview metadata: %v", err)
// 		}

// 		decoder := gob.NewDecoder(reader)
// 		var metadata PreviewMetadata

// 		if err := decoder.Decode(&metadata); err != nil {
// 			return nil, fmt.Errorf("Failed to decode preview metadata: %v", err)
// 		}

// 		if err := reader.Close(); err != nil {
// 			return nil, fmt.Errorf("Failed to close reader: %v", err)
// 		}

// 		name := attr.Prefix[:len(attr.Prefix)-1]
// 		data[name] = metadata
// 	}

// 	return data, nil
// }

// func GetColumns(cacheTreeName string, xColumnNames []string, yColumnNames []string) ([]XColumnNode, []YColumnNode, error) {
// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx)

// 	if err != nil {
// 		return nil, nil, fmt.Errorf("Failed to create GCS client: %v", err)
// 	}

// 	defer client.Close()

// 	bucket := client.Bucket(bucketInfo.BucketName)

// 	xColumnChan := make(chan XColumnNode, len(xColumnNames))
// 	yColumnChan := make(chan YColumnNode, len(yColumnNames))
// 	errorChan := make(chan error)

// 	fetchXColumn := func(name string) {
// 		obj := bucket.Object(cacheTreeName + "/" + xColumnsSubdir + "/" + name)
// 		reader, err := obj.NewReader(ctx)

// 		if err != nil {
// 			errorChan <- fmt.Errorf("Failed to read X column %s: %v", name, err)
// 			return
// 		}

// 		var xCol XColumnNode

// 		{
// 			decoder := gob.NewDecoder(reader)

// 			if err := decoder.Decode(&xCol); err != nil {
// 				errorChan <- fmt.Errorf("Failed to decode X column %s: %v", name, err)
// 				return
// 			}
// 		}

// 		if err := reader.Close(); err != nil {
// 			errorChan <- fmt.Errorf("Failed to close reader for X column %s: %v", name, err)
// 			return
// 		}

// 		xColumnChan <- xCol
// 	}

// 	fetchYColumn := func(name string) {
// 		obj := bucket.Object(cacheTreeName + "/" + yColumnsSubdir + "/" + name)
// 		reader, err := obj.NewReader(ctx)

// 		if err != nil {
// 			errorChan <- fmt.Errorf("Failed to read Y column %s: %v", name, err)
// 			return
// 		}

// 		var yCol YColumnNode

// 		{
// 			decoder := gob.NewDecoder(reader)

// 			if err := decoder.Decode(&yCol); err != nil {
// 				errorChan <- fmt.Errorf("Failed to decode Y column %s: %v", name, err)
// 				return
// 			}
// 		}

// 		if err := reader.Close(); err != nil {
// 			errorChan <- fmt.Errorf("Failed to close reader for Y column %s: %v", name, err)
// 			return
// 		}

// 		yColumnChan <- yCol
// 	}

// 	for _, name := range xColumnNames {
// 		go fetchXColumn(name)
// 	}

// 	for _, name := range yColumnNames {
// 		go fetchYColumn(name)
// 	}

// 	xColumns := make([]XColumnNode, 0, len(xColumnNames))
// 	yColumns := make([]YColumnNode, 0, len(yColumnNames))

// 	select {
// 	case xCol := <-xColumnChan:
// 		xColumns = append(xColumns, xCol)
// 	case yCol := <-yColumnChan:
// 		yColumns = append(yColumns, yCol)
// 	case err := <-errorChan:
// 		return nil, nil, err
// 	}

// 	return xColumns, yColumns, nil
// }

// func DeleteCache(name string) error {
// 	return nil
// }
