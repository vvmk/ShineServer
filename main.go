package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

const port = ":8080"

var globalSessions *session.Manager

func init() {
	globalSessions = NewManager("memory", "gosessionid", 3600)
}

func main() {
	router := NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(router)))
}
