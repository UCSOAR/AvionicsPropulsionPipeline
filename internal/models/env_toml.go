package models

import set "soarpipeline/pkg/set"

type EnvToml struct {
	GoogleClientID     string   `toml:"google_client_id"`
	GoogleClientSecret string   `toml:"google_client_secret"`
	SigningKey         string   `toml:"signing_key"`
	InProduction       bool     `toml:"in_production"`
	Whitelist          []string `toml:"whitelist"`

	Dev struct {
		Host string   `toml:"host"`
		Port string   `toml:"port"`
		Cors []string `toml:"cors"`
	} `toml:"dev"`

	Prod struct {
		Host string   `toml:"host"`
		Port string   `toml:"port"`
		Cors []string `toml:"cors"`
	} `toml:"prod"`
}

func (e *EnvToml) ToAppConfig() AppConfig {
	whitelistSet := set.NewHashSet[string]()

	for _, item := range e.Whitelist {
		whitelistSet.Put(item)
	}

	// Choose host based on environment
	host := e.Dev.Host
	port := e.Dev.Port
	cors := e.Dev.Cors
	if e.InProduction {
		host = e.Prod.Host
		port = e.Prod.Port
		cors = e.Prod.Cors
	}

	config := AppConfig{
		InProduction: e.InProduction,
		SigningKey:   []byte(e.SigningKey),
		Whitelist:    whitelistSet,
		Host:         host,
		Port:         port,
		Cors:         cors,
	}

	return config
}
