package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/goware/emailx"
	"github.com/sendgrid/rest"
	"github.com/vvmk/bounce/mail"
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

type NewUserRequest struct {
	Email    string
	Password string
	Tag      string
	Main     string
}

func Register(w http.ResponseWriter, r *http.Request) {
	var nu NewUserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&nu)
	if err != nil {
		panic(err)
	}

	email := emailx.Normalize(nu.Email)

	// looks better than nil-ing err
	e := emailx.ValidateFast(email)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid email"))
		log.Printf("invalid email: %s", email)
		return
	}

	password, err := HashPassword(nu.Password)
	if err != nil {
		panic(err)
	}

	user := &models.User{
		Email:     email,
		Confirmed: false,
		Hash:      password,
		Tag:       nu.Tag,
		Main:      nu.Main,
		Bio:       "",
	}

	id, err := env.db.CreateUser(user)
	if err != nil {
		panic(err)
	}

	// TODO: I might want the mail pkg to be responsible for its responses
	mailerResponse, err := mail.SendConfirmation(user.Tag, user.Email)
	if err != nil {
		panic(err)
	}

	res := struct {
		user_id         int
		message         string
		mailer_response *rest.Response
	}{
		user_id:         id,
		message:         "Confirmation email sent",
		mailer_response: mailerResponse,
	}

	// TODO: 200 and email sent/received but empty object in response body
	// either respond with a redirect to a thank you page or the new user profile
	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

// TODO: ConfirmUser handles a user clicking an emailed confirmation link
// and updates their db entry to confirmed=true
func ConfirmUser(w http.ResponseWriter, r *http.Request) {

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

func EditUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
