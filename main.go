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
	handler := func(s *server) {
		// Wrap the server's routes with a logging middleware
		s.handler = handlers.LoggingHandler(out, s.routes())
	}
	srv := newServer(
		withAddr(getAddr()),
		handler,
	)
	return start(srv)
}

func getAddr() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return ":8080"
	}
	return net.JoinHostPort("", port)
}
