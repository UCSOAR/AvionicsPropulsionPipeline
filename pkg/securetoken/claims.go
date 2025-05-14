package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
)

const (
	issuer = "SOAR"
)

var (
	errMalformedToken = errors.New("malformed token")
)

func MakeRegisteredClaims(expiry time.Duration) jwt.RegisteredClaims {
	issuedAt := jwt.NewNumericDate(time.Now())
	expiresAt := jwt.NewNumericDate(issuedAt.Add(expiry))
	claims := jwt.RegisteredClaims{
		Issuer:    issuer,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}

	return claims
}

func SignClaims[C jwt.Claims](claims C, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)

	return signedToken, err
}

func ExtractClaims[C jwt.Claims](signedString string, secret []byte) (C, error) {
	var claims C

	token, err := jwt.Parse(signedString, func(_ *jwt.Token) (any, error) {
		return secret, nil
	})

	if err != nil {
		return claims, err
	}

	claimsMap, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return claims, errMalformedToken
	}

	if err := mapstructure.Decode(claimsMap, &claims); err != nil {
		return claims, err
	}

	return claims, nil
}
