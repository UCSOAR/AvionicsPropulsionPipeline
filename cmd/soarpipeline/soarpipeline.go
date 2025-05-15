package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	controllers "soarpipeline/internal/controllers"
	middlewares "soarpipeline/internal/middlewares"
	storage "soarpipeline/internal/storage"
)

const (
	addr         = ":8080"
	readTimeout  = 10 * time.Second
	writeTimeout = 15 * time.Second
	idleTimeout  = 10 * time.Second
)

var (
	errMissingClientID     = errors.New("missing client ID")
	errMissingClientSecret = errors.New("missing client secret")
	errMissingInProduction = errors.New("missing in production flag")
	errMissingSigningKey   = errors.New("missing signing key")
)

func initDependencyInjection() (*controllers.DependencyInjection, error) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")

	if len(clientID) == 0 {
		return nil, errMissingClientID
	}

	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	if len(clientSecret) == 0 {
		return nil, errMissingClientSecret
	}

	// This should match route for callback in the router
	redirectURL := fmt.Sprintf("http://localhost%s/auth/google/callback", addr)

	oauthCfg := oauth2.Config{
		RedirectURL:  redirectURL,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	inProduction, err := strconv.ParseBool(os.Getenv("IN_PRODUCTION"))

	if err != nil {
		return nil, errMissingInProduction
	}

	signingKey := []byte(os.Getenv("SIGNING_KEY"))

	if len(signingKey) == 0 {
		return nil, errMissingSigningKey
	}

	injection := &controllers.DependencyInjection{
		OAuthCfg:     oauthCfg,
		InProduction: inProduction,
		SigningKey:   signingKey,
	}

	// NOTE: dependency injection pointer escapes to heap here
	return injection, nil
}

func main() {
	// Ensure storage directories are initialized
	if err := storage.InitStorageDirectories(); err != nil {
		panic(err)
	}

	// Read .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	injection, err := initDependencyInjection()

	if err != nil {
		panic(err)
	}

	// Set up the router and middleware
	r := chi.NewRouter()
	middlewares.UseCorsMiddleware(r, injection.InProduction)

	// Subrouter for authentication
	r.Route("/auth", func(r chi.Router) {
		r.Get("/me", injection.GetMe)
		r.Post("/logout", injection.PostLogout)

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
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
