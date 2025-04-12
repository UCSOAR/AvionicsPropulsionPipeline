package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// MockServer function to simulate API responses
func mockServer(handler http.Handler) *httptest.Server {
	return httptest.NewServer(handler)
}
