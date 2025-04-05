package displayusage

import (
	"net/http"
	"testing"
)

// StorageResponse struct to hold the API response
type StorageResponse struct {
	StorageSizeBytes int64 `json:"storageSizeBytes"`
}


// Test getStorageUsage with a valid response
func TestGetStorageUsage_Success(t *testing.T) {
	mockResp := StorageResponse{StorageSizeBytes: 1048576} // 1MB
	server := mockServer(mockResp, http.StatusOK)
	defer server.Close()

	storageSize, err := getStorageUsage(server.URL)
	if err != nil {
		t.Errorf("getStorageUsage() returned error: %v", err)
	}

	expected := mockResp.StorageSizeBytes
	if storageSize != expected {
		t.Errorf("Expected storage size %d, got %d", expected, storageSize)
	}
}

// Test API failure scenario
func TestGetStorageUsage_ErrorResponse(t *testing.T) {
	server := mockServer(StorageResponse{}, http.StatusInternalServerError)
	defer server.Close()

	_, err := getStorageUsage(server.URL)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
