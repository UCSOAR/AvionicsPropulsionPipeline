package controllers

import (
	"net/http"
	"time"
)

func (d *DependencyInjection) PostLogout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     SessionCookieName,
		Value:    "",
		Path:     sessionCookiePath,
		HttpOnly: true,
		Secure:   d.AppConfig.InProduction,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}
