package main

import (
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

// GetJWT returns a new JWT token for an authenticated user
func GetJWT(u *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": false,
		"uid":   u.UserId,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(ssr_jwt_key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// jwtMiddleware checks for a valid JWT on protected routes
var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return ssr_jwt_key, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
