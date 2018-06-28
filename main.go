package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/namsral/flag"
)

const port = ":8080"

var ssr_jwt_key []byte

func main() {

	var ssr_jwt string
	flag.StringVar(&ssr_jwt, "ssrjwt", "", "ssr jwt signing key")
	flag.Parse()

	ssr_jwt_key = []byte(ssr_jwt)

	router := NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(router)))
}
