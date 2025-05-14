package token

import "github.com/golang-jwt/jwt/v5"

type GoogleUserClaims struct {
	jwt.RegisteredClaims
	Email         string `json:"email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}
