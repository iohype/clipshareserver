package main

import (
	"log"
	"net/http"
	"os"
)

type option func(*server)

type server struct {
	Addr    string
	handler http.Handler
	db      DB
	logger  *log.Logger
}

//newServer returns a server after applying all options
func newServer(opts ...option) *server {
	srv := &server{}
	// Init defaults
	srv.handler = srv.routes()
	srv.logger = log.New(os.Stdout, "", log.LstdFlags)
	// Apply all options
	for _, opt := range opts {
		opt(srv)
	}
	return srv
}

func start(srv *server) error {
	log.Printf("Starting server on %s", srv.Addr)
	err := http.ListenAndServe(srv.Addr, srv)
	if err != nil {
		return err
	}
	return nil
}

func withLogger(l *log.Logger) option {
	return func(srv *server) {
		srv.logger = l
	}
}

func withAddr(addr string) option {
	return func(srv *server) {
		srv.Addr = addr
	}
}

func withHandler(h http.Handler) option {
	return func(srv *server) {
		srv.handler = h
	}
}

func withDB(db DB) option {
	return func(srv *server) {
		srv.db	= db
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}
