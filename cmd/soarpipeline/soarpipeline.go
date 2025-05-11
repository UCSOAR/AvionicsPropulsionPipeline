package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	controllers "soarpipeline/internal/controllers"
	middlewares "soarpipeline/internal/middlewares"
	storage "soarpipeline/internal/storage"
)

const addr = ":8080"

func main() {
	// Ensure storage directories are initialized
	if err := storage.InitStorageDirectories(); err != nil {
		panic(err)
	}

	// Read .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	oauthCfg := &oauth2.Config{
		RedirectURL:  fmt.Sprintf("http://localhost%s/auth/google/callback", addr),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	injection := controllers.DependencyInjection{
		OAuthCfg:          oauthCfg,
		RandomStringState: os.Getenv("RANDOM_STRING_STATE"),
	}

	// Set up the router and middleware
	r := chi.NewRouter()
	middlewares.UseCorsMiddleware(r)

	// Subrouter for authentication
	r.Route("/auth", func(r chi.Router) {

		// Subrouter for Google OAuth
		r.Route("/google", func(r chi.Router) {
			r.Get("/login", injection.GetGoogleLogin)
			r.Get("/callback", injection.GetGoogleCallback) // This should match the redirect URL in the OAuth config
		})
	})

	// Subrouter for API
	r.Route("/api", func(r chi.Router) {
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
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
