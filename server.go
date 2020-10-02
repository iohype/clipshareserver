package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type server struct{
	Addr string
	router *mux.Router
}

func start(srv *server) error {
	srv.setupRoutes()
	log.Printf("Starting server on %s", srv.Addr)
	err := http.ListenAndServe(srv.Addr, srv.router)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}