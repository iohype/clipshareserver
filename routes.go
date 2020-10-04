package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *server) routes() http.Handler {
	r := mux.NewRouter()
	// Register paths
	// GET /clips
	r.HandleFunc("/clips", s.handleAuthed(s.handleClipsGet())).Methods(http.MethodGet)
	return r
}