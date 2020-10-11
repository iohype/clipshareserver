package main

import (
	"log"
	"net/http"
	"os"
)

type option func(*server)

type server struct {
	// Addr is the address to run the server on
	Addr     string

	handler  http.Handler
	db       DB
	logger   *log.Logger
	verifier verifier
}

func newServer(opts ...option) (*server, error) {
	srv := &server{}
	srv.handler = srv.routes()
	srv.logger = log.New(os.Stdout, "", log.LstdFlags)
	srv.db = newInMemDB()

	for _, opt := range opts {
		opt(srv)
	}

	//error-able defaults are set here to allow users provide alternate impls.
	//They will not be used if the user provides an alternative.
	var err error
	if srv.verifier == nil {
		srv.verifier, err = newFirebaseVerifier()
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
