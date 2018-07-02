package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/goware/emailx"
	"github.com/vvmk/bounce/models"
)

const JSON = "application/json; charset=UTF-8"

func Login(w http.ResponseWriter, r *http.Request) {

	// TODO: email and ok? omitted until the db exists
	_, password, _ := r.BasicAuth()

	// TODO: if email not found, return 401.
	hash, _ := HashPassword("secret")
	user := RepoFindUser(1)

	if !CheckPasswordHash(password, hash) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", JSON)

	token, err := GetJWT(&user)
	if err != nil {
		panic(err)
	}

	t := struct {
		Token string `json:"token"`
	}{token}

	json.NewEncoder(w).Encode(t)
}

func Register(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := emailx.Normalize(vars["email"])

	err := emailx.ValidateFast(email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password, err := HashPassword(vars["password"])
	if err != nil {
		panic(err)
	}

	user := &models.User{
		Email:     email,
		Confirmed: false,
		Hash:      password,
		Tag:       vars["tag"],
		Main:      vars["main"],
		Bio:       "",
	}

	id, err := env.db.CreateUser(user)
	if err != nil {
		panic(err)
	}

	// TODO: Dispatch confirmation email and redirect to Login

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"id\":\"%d\"}", id)
}

func GetLibrary(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user")

	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		// TODO: get the claims and check for admin, uid, expiration, etc...
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

func CreateRoutine(w http.ResponseWriter, r *http.Request) {
}

func ForkRoutine(w http.ResponseWriter, r *http.Request) {
}

func EditRoutine(w http.ResponseWriter, r *http.Request) {
}

func DeleteRoutine(w http.ResponseWriter, r *http.Request) {
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
}

func EditUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
