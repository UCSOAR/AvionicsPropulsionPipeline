package auth

import (
	"log"
	"os"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

// Define a custom type for the key to avoid collisions
type ContextKey string

const (
	// Create a unique key for the provider in context
	ProviderKey ContextKey = "provider"
	key         = "something-random-and-secure"
	MaxAge      = 86400 * 30
	IsProd      = false
)

func NewAuth() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	// Initialize session store
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)
	store.Options = &sessions.Options{
		Path:     "/",
		Domain:   "localhost", // 👈 Required for cookies to work across ports on localhost
		HttpOnly: true,
		Secure:   IsProd,      // false for local dev, true for production
		SameSite: http.SameSiteLaxMode, // 👈 This is essential for cross-origin redirects
	}

	// Configure Gothic
	gothic.Store = store

	// Register Google provider with Gothic
	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, "http://localhost:8080/auth/google/callback"),
	)
}

