package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"OAuth/internal/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Your frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)
	r.Get("/api/user", s.userHandler)

	// OAuth routes
	r.Get("/auth/{provider}", gothic.BeginAuthHandler)
	r.Get("/auth/{provider}/callback", s.getAuthCallbackFunction)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello World"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("JSON marshal error: %v", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (s *Server) getAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	// Store provider value in context using the custom key defined in auth package
	r = r.WithContext(context.WithValue(r.Context(), auth.ProviderKey, provider))

	// Complete the authentication using Gothic
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Printf("OAuth callback error: %v", err)
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}

	log.Printf("Authenticated user: %+v\n", user)

	// Save user to session so frontend can access it later
	session, _ := gothic.Store.Get(r, "gothic-session")
	session.Values["user"] = user
	session.Save(r, w)

	// Redirect user to frontend
	http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
}

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := gothic.Store.Get(r, "gothic-session")
	user, ok := session.Values["user"]
	if !ok {
		http.Error(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
