package utils

import (
	"errors"
	"os"
	"strings"
)

// keep the string after the "TOKEN-" part of the string
func StripBearerPrefixFromTokenString(token string) (string, error) {
	if len(token) > 5 && strings.ToUpper(token[0:6]) == "TOKEN" {
		return token[6:], nil
	} else {
		return token, errors.New("Token is Invalid")
	}
}

func GetJWTSecret() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "just in case"
	}
	return secretKey
}