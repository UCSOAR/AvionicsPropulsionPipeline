package controllers

import "net/http"

func GetStaticFireMetadata(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetStaticFireMetadata"))
}
