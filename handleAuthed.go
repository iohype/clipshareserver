package main

import (
	"net/http"
)

//handleAuthed is a middleware that allows the request only if user is logged in
func (s *server) handleAuthed(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		//next(w, r)
	}
}