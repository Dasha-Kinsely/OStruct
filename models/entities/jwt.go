package entities

import "github.com/golang-jwt/jwt"

// as required by the golang-jwt package
type JWTCustomClaim struct {
	Uid string `json:"user_id"`
	jwt.StandardClaims
}