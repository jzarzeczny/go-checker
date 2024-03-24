package validator

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func ValidateToken(r *http.Request, token string) error {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return errors.New("authorization header is missing")
	}

	parts := strings.Split(authHeader, " ")

	fmt.Println(parts)
	fmt.Println(token)

	if len(parts) != 2 || parts[0] != "Bearer" {
		return errors.New("invalid Authorization header format")
	}

	return nil
}
