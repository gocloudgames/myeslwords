package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorizationss")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	const prefix = "ApiKey "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.New("authorization header must start with 'ApiKey '")
	}

	key := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))
	if key == "" {
		return "", errors.New("token key is empty")
	}

	return key, nil
}
