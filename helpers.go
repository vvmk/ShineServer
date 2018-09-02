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

// MakeRoutineHeaders scrubs a slice of Routines into a slice of RoutineHeaders
// RoutineHeaders consist only of data needed by a user's library page.
func MakeRoutineHeaders(rs []*models.Routine) []RoutineHeader {
	rsm := make([]RoutineHeader, len(rs))
	for i, r := range rs {
		rsm[i] = RoutineHeader{
			r.RoutineId,
			r.Title,
			r.TotalDuration,
			r.Character,
			r.Popularity,
			r.Description,
		}
	}
	return rsm
}
