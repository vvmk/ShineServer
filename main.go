package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vvmk/shineserver/models"

	"github.com/gorilla/handlers"
)

const port = ":8080"

type Env struct {
	db models.Datastore
}

var env Env

var ssr_jwt_key []byte

func main() {

	// set log flags
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ssr_jwt_key = []byte(os.Getenv("SSRJWT"))

	// init db
	dbU := os.Getenv("POSTGRES_USER")
	dbP := os.Getenv("POSTGRES_PASSWORD")
	dbD := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbU, dbP, dbD)
	db, err := models.NewDB(connStr)
	if err != nil {
		panic(err)
	}

	// inject
	env = Env{db}

	// serve
	router := NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(router)))
}
