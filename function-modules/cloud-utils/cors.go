package cloudUtils

import (
	"net/http"
	"strings"
)

type Cors struct {
	AllowOrigin  string
	AllowMethods []string
	AllowHeaders []string
}

func SetCorsHeaders(w http.ResponseWriter, cors Cors) {
	w.Header().Set("Access-Control-Allow-Origin", cors.AllowOrigin)
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(cors.AllowMethods, ", "))
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(cors.AllowHeaders, ", "))
}
