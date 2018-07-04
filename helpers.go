package main

import (
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func UserAuthorized(r *http.Request) bool {
	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])
	tokenData := r.Context().Value("user")
	claims := tokenData.(*jwt.Token).Claims.(jwt.MapClaims)

	return (claims["admin"].(bool) || userId == claims["uid"].(int))
}
