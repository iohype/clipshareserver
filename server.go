package main

import (
	"github.com/gorilla/mux"
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