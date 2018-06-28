package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const JSON = "application/json; charset=UTF-8"

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping, %q", html.EscapeString(r.URL.Path))
}

func GetRoutine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	routineId, err := strconv.Atoi(vars["routineId"])
	if err != nil {
		panic(err)
	}

	routine := RepoFindRoutine(routineId)

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(routine); err != nil {
		panic(err)
	}
}

func GetLibrary(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user")
	fmt.Fprintf(w, "this is an authenticated request")
	fmt.Fprintf(w, "Claim content:\n")
	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		fmt.Fprintf(w, "%s : \t%#v\n", k, v)
	}

	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		panic(err)
	}

	library := RepoFindLibrary(userId)

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(library); err != nil {
		panic(err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		panic(err)
	}

	user := RepoFindUser(userId)

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	creds := map[string]string{"vvmk": "secret"}

	tag := r.Header.Get("tag")
	pass := r.Header.Get("pass")

	// login is good, give back a token
	if creds[tag] == pass {
		w.Header().Set("Content-Type", JSON)
		//w.Header().Set("token", token)

		token, err := GetJWT()
		if err != nil {
			panic(err)
		}

		t := struct {
			Token string `json:"token"`
		}{token}

		json.NewEncoder(w).Encode(t)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
