package securetokentest_test

import (
	"testing"
	"time"

	securetoken "soarpipeline/pkg/securetoken"

	"github.com/golang-jwt/jwt/v5"
)

type testClaims[T any] struct {
	jwt.RegisteredClaims
	Data T `json:"data"`
}

func TestSignAndExtractClaims(t *testing.T) {
	secret := []byte("test_secret")
	claims := testClaims[string]{
		RegisteredClaims: securetoken.MakeRegisteredClaims(time.Hour),
		Data:             "test_data",
	}

	signedToken, err := securetoken.SignClaims(claims, secret)

	if err != nil {
		t.Fatalf("Failed to sign claims: %v", err)
	}

	extractedClaims, err := securetoken.ExtractClaims[testClaims[string]](signedToken, secret)

	if err != nil {
		t.Fatalf("Failed to extract claims: %v", err)
	}

	if extractedClaims.Data != claims.Data {
		t.Errorf("Extracted claims data mismatch: got %v, want %v", extractedClaims.Data, claims.Data)
	}
}

func TestExpiredClaimsFailToExtract(t *testing.T) {
	secret := []byte("test_secret")
	claims := testClaims[string]{
		RegisteredClaims: securetoken.MakeRegisteredClaims(0),
		Data:             "test_data",
	}

	signedToken, err := securetoken.SignClaims(claims, secret)

	if err != nil {
		t.Fatalf("Failed to sign claims: %v", err)
	}

	// Simulate token expiration
	time.Sleep(1 * time.Second)
	_, err = securetoken.ExtractClaims[testClaims[string]](signedToken, secret)

	if err == nil {
		t.Fatalf("Expected error when extracting expired claims, got nil")
	}
}
