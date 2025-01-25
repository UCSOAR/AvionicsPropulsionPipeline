package cloudUtils

import (
	"net/http"
	"strings"
)

// Represents the CORS configuration for an HTTP response.
// It contains the allowed origin, methods, and headers.
type Cors struct {
	AllowOrigin  string
	AllowMethods []string
	AllowHeaders []string
}

// Helper function to set CORS headers on an HTTP response.
func SetCorsHeaders(w http.ResponseWriter, cors Cors) {
	w.Header().Set("Access-Control-Allow-Origin", cors.AllowOrigin)
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(cors.AllowMethods, ", "))
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(cors.AllowHeaders, ", "))
}
