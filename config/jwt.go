package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("asd821hda9sbasdk923rnsfs")

type JWTClaim struct {
	UserId int64
	Email string
	jwt.RegisteredClaims
}