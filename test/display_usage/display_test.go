package displayusage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// StorageResponse struct to hold the API response
type StorageResponse struct {
	StorageSizeBytes int64 `json:"storageSizeBytes"`
}

// getStorageUsage function to fetch storage size from the API
func getStorageUsage(url string) (int64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API error: status code %d", resp.StatusCode)
	}

	var storageResp StorageResponse
	if err := json.NewDecoder(resp.Body).Decode(&storageResp); err != nil {
		return 0, err
	}

	return storageResp.StorageSizeBytes, nil
}

// mockServer function to simulate API responses
func mockServer(response StorageResponse, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response)
	})
	return httptest.NewServer(handler)
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
