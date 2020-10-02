package main

import (
	"github.com/gorilla/mux"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	srv := &server{
		Addr: ":8080",
		router: mux.NewRouter(),
	}
	err := start(srv)
	if err != nil {
		return err
	}
	return nil
}