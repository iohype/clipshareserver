package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *server) routes() http.Handler {
	r := mux.NewRouter()

	// Register paths
	// GET /clips
	r.HandleFunc("/clips", s.requireAuthed(s.handleClipsGet())).Methods(http.MethodGet)
	// POST /clips
	r.HandleFunc("/clips", s.requireAuthed(s.handleClipsCreate())).Methods(http.MethodPost)

	return r
}
