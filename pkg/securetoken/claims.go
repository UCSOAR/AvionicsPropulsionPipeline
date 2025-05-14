package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func ExtractClaims[C jwt.Claims](signedString string, secret []byte) (*C, error) {
	claims := new(C)

	token, err := jwt.ParseWithClaims(signedString, *claims, func(token *jwt.Token) (any, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if decodedClaims, ok := token.Claims.(C); ok && token.Valid {
		*claims = decodedClaims
		return claims, nil
	}

	return nil, errMalformedToken
}
