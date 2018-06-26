package main

import "net/http"

type auth struct {
	wrappedHandler http.Handler
}

func Authenticate(h http.Handler) auth {
	return auth{h}
}
