package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	controllers "soarpipeline/internal/controllers"
	middlewares "soarpipeline/internal/middlewares"
	storage "soarpipeline/internal/storage"
)

const addr = ":8080"

func main() {
	// Read .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Ensure storage directories are initialized
	if err := storage.InitStorageDirectories(); err != nil {
		panic(err)
	}

	// Set up the router
	router := chi.NewRouter()
	middlewares.UseCorsMiddleware(router)

	// Subrouter for API
	router.Route("/api", func(r chi.Router) {
		r.Get("/usage", controllers.GetStorageUsage)

		// Subrouter for static fire data
		r.Route("/staticfire", func(r chi.Router) {
			r.Get("/metadata", controllers.GetStaticFireMetadata)
			r.Post("/columns", controllers.PostStaticFireColumns)
			r.Post("/upload", controllers.PostUploadStaticFire)
		})
	})

	fmt.Println("Server running on http://localhost" + addr)

	// Start the server
	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
