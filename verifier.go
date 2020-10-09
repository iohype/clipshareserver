package main

type verifier interface {
	verify(idToken string) (userID string, err error)
}
