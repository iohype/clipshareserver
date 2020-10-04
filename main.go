package main

import (
	"log"
	"net"
	"os"
	"github.com/gorilla/handlers"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	srv := &server{
		Addr: getAddr(),
	}
	// Wrap the server's router with a logging middleware
	srv.handler = handlers.LoggingHandler(os.Stdout, srv.getRouter())
	return start(srv)
}

func getAddr() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return ":8080"
	}
	return net.JoinHostPort("", port)
}