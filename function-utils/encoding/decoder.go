package encoding

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func BinaryDecode[T any](metadata []byte) (T, error) {
	var decodedMetadata T

	buffer := bytes.NewBuffer(metadata)
	decoder := gob.NewDecoder(buffer)

	// Decode the metadata
	if err := decoder.Decode(&decodedMetadata); err != nil {
		return decodedMetadata, fmt.Errorf("Failed to binary decode metadata: %v", err)
	}

	return decodedMetadata, nil
}
