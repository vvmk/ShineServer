package main

import (
	"log"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/vvmk/shineserver/models"
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

func MakeRoutineHeaders(rs []*models.Routine) []RoutineHeader {
	rsm := make([]RoutineHeader, len(rs))
	for i, r := range rs {
		rsm[i] = RoutineHeader{
			r.RoutineId,
			r.Title,
			r.TotalDuration,
			r.Popularity,
			r.Description,
		}
	}
	return rsm
}
