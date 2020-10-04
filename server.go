package main

import (
	"log"
	"net/http"
)

type option func(*server) error

type server struct {
	Addr    string
	handler http.Handler
	db      DB
}

//newServer returns a server after applying all options
func newServer(opts ...option) (*server, error) {
	srv := &server{}
	for _, opt := range opts {
		err := opt(srv)
		if err != nil {
			return srv, err
		}
	}
	return srv, nil
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
