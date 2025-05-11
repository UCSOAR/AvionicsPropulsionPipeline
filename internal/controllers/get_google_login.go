package controllers

import (
	"net/http"

	"golang.org/x/oauth2"
)

func (d *DependencyInjection) GetGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := d.OAuthCfg.AuthCodeURL(d.RandomStringState, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
