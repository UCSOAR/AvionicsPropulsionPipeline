package controllers

import (
	"encoding/json"
	"net/http"
	securetoken "soarpipeline/pkg/securetoken"
)

func (d *DependencyInjection) GetMe(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(sessionCookieName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	signedString := cookie.Value
	claims, err := securetoken.ExtractClaims[securetoken.GoogleUserClaims](signedString, d.SigningKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(claims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
