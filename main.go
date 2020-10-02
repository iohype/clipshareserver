package main

import (
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	srv := newServer()
	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		return err
	}
	return nil
}