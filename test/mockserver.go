package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// MockServer function to simulate API responses
func MockServer(response interface{}, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response)
	})
	return httptest.NewServer(handler)
}
