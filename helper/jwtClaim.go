package helper

import (
	"net/http"

	"github.com/FireCrackerTeam/wonderjack-web/config"
	"github.com/golang-jwt/jwt/v4"
)

func JWTClaim(r *http.Request, token string) (int64, error){
	// Retrieve the JWT token from the request header or form data
	// c, err := r.Cookie("token")
	// 	if err != nil {
	// 		return -1, err
	// 	}

	// println(tokenReq.Token)

	// // get token value
	tokenString := token

	claims := &config.JWTClaim{}
	// parsing token jwt
	jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})

	userID := claims.UserId

	return userID, nil
}
