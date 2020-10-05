package main

import "context"

//verifier verifies a token and returns a uid
type verifier interface {
	//verify a token exists and return a uid for it
	verify(context.Context, string) (string, error)
}