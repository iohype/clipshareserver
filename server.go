package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type server struct{
	router *mux.Router
}

func newServer() *server {
	srv := &server{}
	srv.router = mux.NewRouter()
	srv.setupRoutes(srv.router)
	return srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}