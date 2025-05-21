package controllers

import (
	"encoding/json"
	"net/http"
	"soarpipeline/internal/models"
	securetoken "soarpipeline/pkg/securetoken"
)

func (d *DependencyInjection) GetMe(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(SessionCookieName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	signedString := cookie.Value
	userClaims, err := securetoken.ExtractClaims[models.GoogleUserClaims](signedString, d.AppConfig.SigningKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	clientUser := userClaims.ToClientUser()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(clientUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
