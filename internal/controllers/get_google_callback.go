package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"soarpipeline/internal/models"
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
	cookieName   = "session"
	cookiePath   = "/"
	cookieMaxAge = 2 * time.Hour
)

var (
	errMissingCode  = errors.New("missing code")
	errMissingState = errors.New("missing state")
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
	token, err := d.OAuthCfg.Exchange(ctx, code)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Use the access token to get user info
	client := d.OAuthCfg.Client(ctx, token)
	res, err := client.Get(oauth2UserInfoEndpoint)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()
	var user models.GoogleUser

	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create cookie with user info
	maxAge := int(cookieMaxAge.Abs().Seconds())
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    user.Email,
		Path:     cookiePath,
		HttpOnly: true,
		Secure:   d.InProduction,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   maxAge,
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, redirectState, http.StatusTemporaryRedirect)
}
