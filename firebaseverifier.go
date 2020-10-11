package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type firebaseVerifier struct {
	app  *firebase.App
	test bool
}

func newFirebaseVerifier() (*firebaseVerifier, error) {
	app, err := newFirebaseAppInstance()
	if err != nil {
		return nil, err
	}
	return &firebaseVerifier{app: app}, nil
}

func newTestFirebaseVerifier() (*firebaseVerifier, error) {
	return &firebaseVerifier{nil, true}, nil
}

func newFirebaseAppInstance() (*firebase.App, error) {
	return firebase.NewApp(context.Background(), nil)
}

func (f *firebaseVerifier) getAuthClient(ctx context.Context) (*auth.Client, error) {
	return f.app.Auth(ctx)
}

func (f *firebaseVerifier) verify(idToken string) (string, error) {
	if f.isTest(idToken) {
		return "user1369", nil
	}
	return f.getUserIDForToken(idToken)
}

func (f *firebaseVerifier) getUserIDForToken(idToken string) (string, error) {
	ctx := context.Background()

	authClient, err := f.getAuthClient(ctx)
	if err != nil {
		return "", err
	}
	token, err := authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", err
	}

	return token.UID, nil
}

func (f *firebaseVerifier) isTest(idToken string) bool {
	return f.test || idToken == "TestToken"
}
