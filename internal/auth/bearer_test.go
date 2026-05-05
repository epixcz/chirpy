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

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey THE_KEY_HERE")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("GetAPIKey returned error: %v", err)
	}

	if apiKey != "THE_KEY_HERE" {
		t.Fatalf("got api key %q, want %q", apiKey, "THE_KEY_HERE")
	}
}

func TestGetAPIKeyTrimsWhitespace(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey   THE_KEY_HERE  ")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("GetAPIKey returned error: %v", err)
	}

	if apiKey != "THE_KEY_HERE" {
		t.Fatalf("got api key %q, want %q", apiKey, "THE_KEY_HERE")
	}
}

func TestGetAPIKeyMissingAuthorizationHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	if err == nil {
		t.Fatal("GetAPIKey accepted missing Authorization header")
	}
}

func TestGetAPIKeyRejectsMalformedAuthorizationHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer THE_KEY_HERE")

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("GetAPIKey accepted malformed Authorization header")
	}
}
