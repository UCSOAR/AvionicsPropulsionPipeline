package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"soarpipeline/internal/models"
	securetoken "soarpipeline/pkg/securetoken"
	utils "soarpipeline/pkg/utils"
	"time"
)

const (
	oauth2UserInfoEndpoint = "https://www.googleapis.com/oauth2/v1/userinfo"
)

const (
	errorParam = "error"
	codeParam  = "code"
	stateParam = "state"
)

const (
	SessionCookieName = "session_token"
	sessionCookiePath = "/"
	tokenExpiry       = 12 * time.Hour
)

var (
	errMissingCode   = errors.New("missing code")
	errMissingState  = errors.New("missing state")
	errForbiddenUser = errors.New("user not in whitelist")
)

func (d *DependencyInjection) GetGoogleCallback(w http.ResponseWriter, r *http.Request) {
	redirectState := r.URL.Query().Get(stateParam)

	if len(redirectState) == 0 {
		http.Error(w, errMissingState.Error(), http.StatusBadRequest)
		return
	}

	if safeRedirectState, err := url.QueryUnescape(redirectState); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		redirectState = safeRedirectState
	}

	if errorMessage := r.URL.Query().Get(errorParam); len(errorMessage) > 0 {
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get(codeParam)

	if len(code) == 0 {
		http.Error(w, errMissingCode.Error(), http.StatusBadRequest)
		return
	}

	// Exchange the code for an access token
	ctx := r.Context()
	accessToken, err := d.OAuthConfig.Exchange(ctx, code)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Use the access token to get user info
	client := d.OAuthConfig.Client(ctx, accessToken)
	res, err := client.Get(oauth2UserInfoEndpoint)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()
	var userClaims models.GoogleUserClaims

	if err := json.NewDecoder(res.Body).Decode(&userClaims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userClaims.RegisteredClaims = securetoken.MakeRegisteredClaims(tokenExpiry)

	// Ensure user is in whitelist
	if !d.AppConfig.Whitelist.Has(userClaims.Email) {
		http.Error(w, errForbiddenUser.Error(), http.StatusForbidden)
		return
	}

	token, err := securetoken.SignClaims(userClaims, d.AppConfig.SigningKey)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create cookie with user info
	cookie := &http.Cookie{
		Name:     SessionCookieName,
		Value:    token,
		Path:     sessionCookiePath,
		HttpOnly: true,
		Secure:   d.AppConfig.InProduction,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   utils.DurationToSeconds(tokenExpiry),
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, redirectState, http.StatusTemporaryRedirect)
}
