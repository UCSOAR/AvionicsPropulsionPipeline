package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	controllers "soarpipeline/internal/controllers"
	middlewares "soarpipeline/internal/middlewares"
	"soarpipeline/internal/models"
	storage "soarpipeline/internal/storage"
)

const (
	readTimeout  = 10 * time.Second
	writeTimeout = 15 * time.Second
	idleTimeout  = 10 * time.Second
)

const (
	addr        = ":8080"
	envTomlFile = ".env.toml"
)

func initDependencyInjection() (*controllers.DependencyInjection, error) {
	var env models.EnvToml

	if _, err := toml.DecodeFile(envTomlFile, &env); err != nil {
		return nil, err
	}

	// This should match route for callback in the router
	redirectURL := fmt.Sprintf("http://localhost%s/auth/google/callback", addr)

	oauthCfg := oauth2.Config{
		RedirectURL:  redirectURL,
		ClientID:     env.GoogleClientID,
		ClientSecret: env.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	appConfig := env.ToAppConfig()
	injection := controllers.DependencyInjection{
		OAuthConfig: oauthCfg,
		AppConfig:   appConfig,
	}

	// Injection struct escapes to heap here
	return &injection, nil
}

func main() {
	// Ensure storage directories are initialized
	if err := storage.InitStorageDirectories(); err != nil {
		panic(err)
	}

	i, err := initDependencyInjection()

	if err != nil {
		panic(err)
	}

	// Set up the router and middleware
	r := chi.NewRouter()
	middlewares.UseCorsMiddleware(r, i.AppConfig.InProduction)

	// Subrouter for authentication
	r.Route("/auth", func(r chi.Router) {
		r.Get("/me", i.GetMe)
		r.Post("/logout", i.PostLogout)

		// Subrouter for Google OAuth
		r.Route("/google", func(r chi.Router) {
			r.Get("/login", i.GetGoogleLogin)
			r.Get("/callback", i.GetGoogleCallback) // This should match the redirect URL in the OAuth config
		})
	})

	// Subrouter for API
	r.Route("/api", func(r chi.Router) {
		middlewares.UseAuthTokenExtractorMiddleware(r, i.AppConfig.SigningKey)

		r.Get("/usage", controllers.GetStorageUsage)

		// Subrouter for static fire data
		r.Route("/staticfire", func(r chi.Router) {
			r.Get("/metadata", controllers.GetStaticFireMetadata)
			r.Post("/columns", controllers.PostStaticFireColumns)
			r.Post("/upload", controllers.PostUploadStaticFire)
			r.Get("/download", controllers.PostStaticFireDownload)

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
