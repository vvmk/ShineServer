package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/goware/emailx"
	"github.com/sendgrid/rest"
	"github.com/vvmk/shineserver/mail"
	"github.com/vvmk/shineserver/models"
)

const JSON = "application/json; charset=UTF-8"

// TokenResponse is for returning a JWT to a validated user.
type TokenResponse struct {
	Token string       `json:"access_token"`
	User  *models.User `json:"user"`
}

// NewUserRequest is exactly what it sounds like. See Register().
type NewUserRequest struct {
	Email    string
	Password string
	Tag      string
	Main     string
}

// User came from an email link and is confirming or resetting something
type EmailData struct {
	Uid   int
	Token string
}

// Login verifies the user's input credentials and checks them against
// the stored data, returning an access token, 404, or 401
func Login(w http.ResponseWriter, r *http.Request) {

	// get the entered credentials
	email, password, ok := r.BasicAuth()
	if !ok {
		panic(errors.New("Basic auth not ok!"))
	}

	// lookup user by email
	user, err := env.db.FindUserByEmail(email)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check against stored password
	if !CheckPasswordHash(password, user.Hash) {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", JSON)

	// TODO: send back the user data with the token
	token, err := GetJWT(user.UserId)
	if err != nil {
		panic(err)
	}

	if err = json.NewEncoder(w).Encode(&TokenResponse{token, user}); err != nil {
		panic(err)
	}
}

// Register handles the creation of a new user. Right now it also handles
// sanitization/validation of the incoming data, a little too much of
// the email process, and has some extra test data jammed at the end in
// addition to its usual handler duties...
// TODO: dire need of refactoring.
func Register(w http.ResponseWriter, r *http.Request) {
	var nu NewUserRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&nu); err != nil {
		panic(err)
	}

	email := emailx.Normalize(nu.Email)

	if err := emailx.ValidateFast(email); err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

	//TODO: these things should be grouped together somehow,
	// too many points of failure
	id, err := env.db.CreateUser(user)
	if err != nil {
		panic(err)
	}

	token, err := GenerateEmailToken()
	if err != nil {
		panic(err)
	}

	if err = env.db.CreateActivation(id, token); err != nil {
		panic(err)
	}

	// TODO: maybe the mail pkg should be responsible for handling more of this
	messageData := &mail.MessageData{
		Address: user.Email,
		Tag:     user.Tag,
		UserId:  id,
		Token:   token,
	}

	mailerResponse, err := mail.SendConfirmation(messageData)
	if err != nil {
		panic(err)
	}

	// TODO: mail response not needed if pkg mail handles/logs its errors
	res := struct {
		UserId         int            `json:"user_id"`
		Message        string         `json:"message"`
		MailerResponse *rest.Response `json:"mailer_response"`
	}{
		UserId:         id,
		Message:        "Confirmation email sent",
		MailerResponse: mailerResponse,
	}

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

// ConfirmUser handles a POST with data from an account confirmation email.
// This POST is intended to parse the email'd link params and use them to
// confirm a new user.
func ConfirmUser(w http.ResponseWriter, r *http.Request) {

	// get params from post body
	body := &EmailData{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		log.Println("ConfirmUser: malformed json data")
		panic(err)
	}

	token := body.Token
	userId := body.Uid

	// try to validate/confirm user
	if err := env.db.ConfirmUser(userId, token); err != nil {
		log.Println("ConfirmUser: User tried to authenticate from email but failed.")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := env.db.FindUserById(userId)
	if err != nil {
		log.Panic("This should have failed sooner or not at all...")
		panic(err)
	}

	// send back a jwt so user is auto logged in
	jwt, err := GetJWT(userId)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&TokenResponse{jwt, user}); err != nil {
		panic(err)
	}
}

// GetLibrary fetches a user's public facing routines (all of them are public)
func GetLibrary(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])

	routines, err := env.db.FindRoutinesByCreator(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(routines); err != nil {
		panic(err)
	}
}

// GetRoutine fetches a single routine using {routineId} or 404
func GetRoutine(w http.ResponseWriter, r *http.Request) {

	routineId, _ := strconv.Atoi(mux.Vars(r)["routineId"])

	routine, err := env.db.FindRoutineById(routineId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(routine); err != nil {
		panic(err)
	}
}

// CreateRoutine handles an authorized user creating a brand new routine
func CreateRoutine(w http.ResponseWriter, r *http.Request) {

	if !UserAuthorized(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var routine models.Routine

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&routine); err != nil {
		panic(err)
	}

	routineId, err := env.db.CreateRoutine(&routine)
	if err != nil {
		panic(err)
	}

	newRoutine, err := env.db.FindRoutineById(routineId)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(newRoutine); err != nil {
		panic(err)
	}
}

// ForkRoutine allows an authorized user to copy a routine.
// A brand new Routine entry is created for the new user's id,
// but retains the id of the original creator separately.
func ForkRoutine(w http.ResponseWriter, r *http.Request) {

	if !UserAuthorized(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["userId"])
	routineId, _ := strconv.Atoi(vars["routineId"])

	routine, err := env.db.FindRoutineById(routineId)
	if err != nil {
		log.Printf("userid: %d, routineId: %d", userId, routineId)
		panic(err)
	}

	// change creator_id
	routine.CreatorId = userId

	// increment popularity
	routine.Popularity = routine.Popularity + 1

	newId, err := env.db.CreateRoutine(routine)
	if err != nil {
		panic(err)
	}

	newRoutine, err := env.db.FindRoutineById(newId)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(newRoutine); err != nil {
		panic(err)
	}
}

// EditRoutine modifies select fields of a user's routine and returns
// the modified routine.
func EditRoutine(w http.ResponseWriter, r *http.Request) {

	if !UserAuthorized(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	routineId, _ := strconv.Atoi(mux.Vars(r)["routineId"])

	var routine models.Routine

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&routine); err != nil {
		panic(err)
	}

	if err := env.db.UpdateRoutine(routineId, &routine); err != nil {
		panic(err)
	}

	newRoutine, err := env.db.FindRoutineById(routineId)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(newRoutine); err != nil {
		panic(err)
	}
}

// DeleteRoutine just deletes a routine if the authorized user is its
// creator. Any routines this one was forked from are unaffected.
func DeleteRoutine(w http.ResponseWriter, r *http.Request) {

	if !UserAuthorized(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	routineId, _ := strconv.Atoi(mux.Vars(r)["routineId"])

	if err := env.db.DeleteRoutine(routineId); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

// GetUser returns a user's public data or 404
func GetUser(w http.ResponseWriter, r *http.Request) {

	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])

	user, err := env.db.FindUserById(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

// EditUser modifies a user's public data if authorized
func EditUser(w http.ResponseWriter, r *http.Request) {

	if !UserAuthorized(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])

	var user models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		panic(err)
	}

	if err := env.db.UpdateUser(userId, &user); err != nil {
		panic(err)
	}

	updatedUser, err := env.db.FindUserById(userId)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		panic(err)
	}
}

// DeleteUser deletes a user if they are authorized...FOR GOOD.
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	if !UserAuthorized(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userId, _ := strconv.Atoi(mux.Vars(r)["userId"])

	if err := env.db.DeleteUser(userId); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
