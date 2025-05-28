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

func UseCorsMiddleware(router chi.Router, allowedOrigins []string) {
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           utils.DurationToSeconds(preflightCacheMaxAge),
	})

	router.Use(corsConfig.Handler)
}
