package main

import (
	"net/http"

	"github.com/go-chi/chi"

	controllers "soarpipeline/internal/controllers"
)

const devAddr = ":8080"

func main() {
	router := chi.NewRouter()

	// Subrouter for API
	router.Route("/api", func(r chi.Router) {
		// Subrouter for static fire data
		r.Route("/staticfire", func(r chi.Router) {
			r.Get("/metadata", controllers.GetStaticFireMetadata)
			r.Post("/columns", controllers.PostStaticFireColumns)
			r.Post("/", controllers.PostStaticFire)
		})
	})

	println("Starting server on", devAddr)

	// Start the server
	if err := http.ListenAndServe(devAddr, router); err != nil {
		panic(err)
	}
}
