package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	tokenString, ok := strings.CutPrefix(authHeader, "Bearer ")
	if !ok {
		return "", errors.New("malformed authorization header")
	}

	tokenString = strings.TrimSpace(tokenString)
	if tokenString == "" {
		return "", errors.New("missing bearer token")
	}

	return tokenString, nil
}

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	apiKey, ok := strings.CutPrefix(authHeader, "ApiKey ")
	if !ok {
		return "", errors.New("malformed authorization header")
	}

	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return "", errors.New("missing api key")
	}

	return apiKey, nil
}
