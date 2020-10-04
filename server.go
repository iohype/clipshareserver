package main

import (
	"log"
	"net/http"
)

type server struct{
	Addr string
	handler http.Handler
}

func start(srv *server) error {
	log.Printf("Starting server on %s", srv.Addr)
	err := http.ListenAndServe(srv.Addr, srv)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}