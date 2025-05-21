package controllers

import (
	"errors"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

const (
	redirectURIParam = "redirect_uri"
)

var (
	errMissingRedirectURI = errors.New("missing redirect URI")
)

func (d *DependencyInjection) GetGoogleLogin(w http.ResponseWriter, r *http.Request) {
	redirect := r.URL.Query().Get(redirectURIParam)

	if len(redirect) == 0 {
		http.Error(w, errMissingRedirectURI.Error(), http.StatusBadRequest)
		return
	}

	// Store redirect URI in state parameter
	state := url.QueryEscape(redirect)
	url := d.OAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
