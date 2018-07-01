package main

import (
	"log"
	"net/http"

	"github.com/vvmk/bounce/models"

	"github.com/gorilla/handlers"
	"github.com/namsral/flag"
)

const port = ":80"

var ssr_jwt_key []byte

type Env struct {
	db models.Datastore
}

func main() {

	// jwt
	var ssr_jwt string
	flag.StringVar(&ssr_jwt, "ssrjwt", "", "ssr jwt signing key")
	flag.Parse()

	ssr_jwt_key = []byte(ssr_jwt)

	// init db
	connStr := "user=ssr-dev dbname=ssr-db sslmode=verify-full"
	db, err := models.NewDB(connStr)
	if err != nil {
		log.Panic(err)
	}

	// inject
	env := &Env{db}

	// serve
	router := NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(router)))
}
