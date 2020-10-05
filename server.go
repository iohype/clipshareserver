package main

import (
	"log"
	"net/http"
	"os"
)

type option func(*server)

type server struct {
	Addr     string
	handler  http.Handler
	db       DB
	logger   *log.Logger
	verifier verifier
}

//newServer returns a server after applying all options
func newServer(opts ...option) (*server, error) {
	srv := &server{}

	srv.Addr = ":8080"
	srv.handler = srv.routes()
	srv.logger = log.New(os.Stdout, "", log.LstdFlags)
	srv.db = newInMemDB()

	// Apply all options
	for _, opt := range opts {
		opt(srv)
	}

	if srv.verifier == nil {
		// Default verifier is fireauth
		var err error
		srv.verifier, err = newFireAuth()
		if err != nil {
			return nil, err
		}
	}

	return srv, nil
}

func start(srv *server) error {
	log.Printf("Starting server on %s", srv.Addr)
	return http.ListenAndServe(srv.Addr, srv)
}

func withVerifier(vfy verifier) option {
	return func(srv *server) {
		srv.verifier = vfy
	}
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
		srv.db = db
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}
