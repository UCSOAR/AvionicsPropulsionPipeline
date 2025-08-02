package models

import set "soarpipeline/pkg/set"

type HostConfig struct {
	Host           string   `toml:"host"`
	Port           string   `toml:"port"`
	AllowedOrigins []string `toml:"allowedorigins"`
}

type EnvToml struct {
	GoogleClientID     string     `toml:"google_client_id"`
	GoogleClientSecret string     `toml:"google_client_secret"`
	SigningKey         string     `toml:"signing_key"`
	InProduction       bool       `toml:"in_production"`
	Whitelist          []string   `toml:"whitelist"`
	Dev                HostConfig `toml:"dev"`
	Prod               HostConfig `toml:"prod"`
}

func (e *EnvToml) ToAppConfig() AppConfig {
	whitelistSet := set.NewHashSet[string]()

	for _, item := range e.Whitelist {
		whitelistSet.Put(item)
	}

	// Select the appropriate environment config
	env := e.Dev
	if e.InProduction {
		env = e.Prod
	}

	config := AppConfig{
		InProduction:   e.InProduction,
		SigningKey:     []byte(e.SigningKey),
		Whitelist:      whitelistSet,
		Host:           env.Host,
		Port:           env.Port,
		AllowedOrigins: env.AllowedOrigins,
	}

	return config
}
