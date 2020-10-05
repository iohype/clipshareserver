package main

import (
	"github.com/gorilla/handlers"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if err := run(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

//run the app using out as the standard output
func run(out io.Writer) error {
	loggedHandler := func(s *server) {
		// Wrap the server's handler with a logging middleware
		s.handler = handlers.LoggingHandler(out, s.handler)
	}
	srv, err := newServer(
		withAddr(getAddr()),
		loggedHandler,
		withLogger(log.New(out, "", log.LstdFlags)),
	)
	if err != nil {
		return err
	}
	return start(srv)
}

func getAddr() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return ":8080"
	}
	return net.JoinHostPort("", port)
}
