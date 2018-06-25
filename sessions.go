package main

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("some-secret"))

func MyHandler(w http.ResponseWriter, r *http.Request) {

	// Get a session. Get() alwats returns a session, even if empty.
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values
	session.Values["foo"] = "bar"
	session.Values[42] = 43

	// Save it before we write to the response/return from the handler.
	// in production, chech for/handle errors here!
	// Save must be called before writing to the response, otherwise
	// the session cookie will not be sent to the client
	session.Save(r, w)
}
