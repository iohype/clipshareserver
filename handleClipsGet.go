package main

import (
	"fmt"
	"net/http"
)

func (s *server) handleClipsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Error(w, r, fmt.Errorf("not implemented"), http.StatusNotImplemented)
	}
}