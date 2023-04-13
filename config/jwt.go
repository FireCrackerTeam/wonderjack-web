package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("asd821hda9sbasdk923rnsfs")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}