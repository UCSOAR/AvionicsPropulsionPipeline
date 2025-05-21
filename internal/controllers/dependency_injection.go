package controllers

import (
	"soarpipeline/internal/models"

	"golang.org/x/oauth2"
)

type DependencyInjection struct {
	OAuthConfig oauth2.Config
	AppConfig   models.AppConfig
}
