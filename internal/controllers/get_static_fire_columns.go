package controllers

import "net/http"

func GetStaticFireColumns(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetStaticFireColumns"))
}
