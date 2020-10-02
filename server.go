package main

type server struct{}

func newServer() *server {
	srv := &server{}
	return srv
}