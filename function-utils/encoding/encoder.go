package encoding

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

func BinaryEncode[T any](metadata T) ([]byte, error) {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)

	// Encode the metadata
	if err := encoder.Encode(metadata); err != nil {
		return nil, fmt.Errorf("Failed to binary encode metadata: %v", err)
	}

	return buffer.Bytes(), nil
}

func JsonEncode[T any](metadata T) (string, error) {
	json, err := json.Marshal(metadata)

	if err != nil {
		return "", fmt.Errorf("Failed to JSON encode metadata: %v", err)
	}

	return string(json), nil
}
