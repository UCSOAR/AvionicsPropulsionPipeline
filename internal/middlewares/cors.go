package middlewares

import (
	utils "soarpipeline/pkg/utils"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

const (
	preflightCacheMaxAge = 300 * time.Second
)

const (
	devCorsOrigin  = "http://localhost:5173"
	prodCorsOrigin = "https://soarpipeline.com"
)

func UseCorsMiddleware(router chi.Router, inProduction bool) {
	allowedOrigin := devCorsOrigin

	if inProduction {
		allowedOrigin = prodCorsOrigin
	}

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{allowedOrigin},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           utils.DurationToSeconds(preflightCacheMaxAge),
	})

	router.Use(corsConfig.Handler)
}
