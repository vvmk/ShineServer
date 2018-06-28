package main

import (
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

func GetJWT() (string, error) {

	// TODO: gut this
	// create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin":   true,
		"user_id": 1,
		"tag":     "vvmk",
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// sign the token
	tokenString, err := token.SignedString(ssr_jwt_key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return ssr_jwt_key, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
