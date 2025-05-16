package models

import set "soarpipeline/pkg/set"

type EnvToml struct {
	GoogleClientID     string   `toml:"google_client_id"`
	GoogleClientSecret string   `toml:"google_client_secret"`
	SigningKey         string   `toml:"signing_key"`
	InProduction       bool     `toml:"in_production"`
	Whitelist          []string `toml:"whitelist"`
}

func (e *EnvToml) ToAppConfig() AppConfig {
	whitelistSet := set.NewHashSet[string]()

	for _, item := range e.Whitelist {
		whitelistSet.Put(item)
	}

	config := AppConfig{
		InProduction: e.InProduction,
		SigningKey:   []byte(e.SigningKey),
		Whitelist:    whitelistSet,
	}

	return config
}
