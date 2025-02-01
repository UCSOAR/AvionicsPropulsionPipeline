package function

import (
	"encoding/json"
	"net/http"

	caching "github.com/UCSOAR/AvionicsPropulsionPipeline/static-fire/caching"
)

func GetStaticFires(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {
	case "GET":
		{
			// Only allow GET requests
			break
		}
	case "OPTIONS":
		{
			w.WriteHeader(http.StatusOK)
			return
		}
	default:
		{
			http.Error(w, "Only GET requests allowed for fetching static fire uploads", http.StatusMethodNotAllowed)
			return
		}
	}

	// Get the static fires from the cache
	staticFires, err := caching.GetAllCacheMetadata()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the static fires as JSON
	json, err := json.Marshal(staticFires)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON to the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
