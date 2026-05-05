package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer TOKEN_STRING")

	tokenString, err := GetBearerToken(headers)
	if err != nil {
		t.Fatalf("GetBearerToken returned error: %v", err)
	}

	if tokenString != "TOKEN_STRING" {
		t.Fatalf("got token %q, want %q", tokenString, "TOKEN_STRING")
	}
}

func TestGetBearerTokenTrimsWhitespace(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer   TOKEN_STRING  ")

	tokenString, err := GetBearerToken(headers)
	if err != nil {
		t.Fatalf("GetBearerToken returned error: %v", err)
	}

	if tokenString != "TOKEN_STRING" {
		t.Fatalf("got token %q, want %q", tokenString, "TOKEN_STRING")
	}
}

func TestGetBearerTokenMissingAuthorizationHeader(t *testing.T) {
	_, err := GetBearerToken(http.Header{})
	if err == nil {
		t.Fatal("GetBearerToken accepted missing Authorization header")
	}
}

func TestGetBearerTokenRejectsMalformedAuthorizationHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Token TOKEN_STRING")

	_, err := GetBearerToken(headers)
	if err == nil {
		t.Fatal("GetBearerToken accepted malformed Authorization header")
	}
}
