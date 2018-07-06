package main

import (
	"log"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func UserAuthorized(r *http.Request) bool {
	userId, err := strconv.ParseFloat(mux.Vars(r)["userId"], 64)
	if err != nil {
		log.Println(err)
	}
	tokenData := r.Context().Value("user")
	claims := tokenData.(*jwt.Token).Claims.(jwt.MapClaims)

	return (claims["admin"].(bool) || userId == claims["uid"].(float64))
}
