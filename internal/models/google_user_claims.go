package models

import "github.com/golang-jwt/jwt/v5"

type GoogleUserClaims struct {
	jwt.RegisteredClaims
	Email         string `json:"email"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}

func (g *GoogleUserClaims) ToClientUser() ClientUser {
	user := ClientUser{
		Email:   g.Email,
		Picture: g.Picture,
	}

	return user
}
