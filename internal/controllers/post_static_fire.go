package controllers

import "net/http"

func PostStaticFire(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PostStaticFire"))
}
