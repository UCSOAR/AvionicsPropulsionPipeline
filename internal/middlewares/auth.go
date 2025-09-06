package middlewares

import (
	"context"
	"net/http"
	"soarpipeline/internal/controllers"
	"soarpipeline/internal/models"
	securetoken "soarpipeline/pkg/securetoken"

	"github.com/go-chi/chi/v5"
)

type authTokenContextKey string

const authKey authTokenContextKey = "auth_token"

func authTokenExtractor(next http.Handler, signingKey []byte) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(controllers.SessionCookieName)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		signedString := cookie.Value
		userClaims, err := securetoken.ExtractClaims[models.GoogleUserClaims](signedString, signingKey)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Store the user claims in the request context
		ctx := context.WithValue(r.Context(), authKey, userClaims)

		// Next middleware
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(handler)
}

func UseAuthTokenExtractorMiddleware(router chi.Router, signingKey []byte) {
	middleware := func(next http.Handler) http.Handler {
		return authTokenExtractor(next, signingKey)
	}

	router.Use(middleware)
}
