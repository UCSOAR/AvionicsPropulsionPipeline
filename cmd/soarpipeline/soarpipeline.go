package main

import (
	"net/http"

	"github.com/go-chi/chi"

	controllers "soarpipeline/internal/controllers"
	middlewares "soarpipeline/internal/middlewares"
	storage "soarpipeline/internal/storage"
)

const devAddr = ":8080"

func main() {
	// Ensure storage directories are initialized
	if err := storage.InitStorageDirectories(); err != nil {
		panic(err)
	}

	// Set up the router
	router := chi.NewRouter()
	middlewares.UseCorsMiddleware(router)

	// Subrouter for API
	router.Route("/api", func(r chi.Router) {
		// Subrouter for static fire data
		r.Route("/staticfire", func(r chi.Router) {
			r.Get("/metadata", controllers.GetStaticFireMetadata)
			r.Post("/columns", controllers.PostStaticFireColumns)
			r.Post("/upload", controllers.PostUploadStaticFire)
		})
	})

	println("Starting server on", devAddr)

	// Start the server
	if err := http.ListenAndServe(devAddr, router); err != nil {
		panic(err)
	}
}
