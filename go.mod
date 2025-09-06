module soarpipeline

go 1.23.5

require (
	github.com/BurntSushi/toml v1.5.0
	github.com/go-chi/chi/v5 v5.2.1
	github.com/go-chi/cors v1.2.1
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/mitchellh/mapstructure v1.5.0
	golang.org/x/oauth2 v0.30.0
)

require cloud.google.com/go/compute/metadata v0.3.0 // indirect
