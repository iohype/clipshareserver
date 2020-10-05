package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
)

//fireAuth Handles firebase authentication
type fireAuth struct {
	app *firebase.App
}

func newFireAuth() (*fireAuth, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return &fireAuth{app}, nil
}

// Implement the verifier interface
func (f *fireAuth) verify(ctx context.Context, idToken string) (string, error) {
	client, err := f.app.Auth(ctx)
	if err != nil {
		return "", err
	}
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}
	return token.UID, nil
}
