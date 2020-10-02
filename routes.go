package main

import "net/http"

func (s *server) setupRoutes() {
	// GET /clips
	s.router.HandleFunc(
		"/clips",
		s.handleAuthed(s.handleClipsGet()),
	).Methods(http.MethodGet)
}