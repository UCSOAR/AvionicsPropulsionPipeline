package controllers

import "golang.org/x/oauth2"

type DependencyInjection struct {
	OAuthCfg     oauth2.Config
	InProduction bool
	SigningKey   []byte
}
