package main

import (
	"io"
	"log"
	"net"
	"os"
	"github.com/gorilla/handlers"
)

func main() {
	if err := run(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

//run the app using out as the standard output
func run(out io.Writer) error {
	srv := &server{
		Addr: getAddr(),
	}
	// Wrap the server's routes with a logging middleware
	srv.handler = handlers.LoggingHandler(out, srv.routes())
	return start(srv)
}

func getAddr() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return ":8080"
	}
	return net.JoinHostPort("", port)
}