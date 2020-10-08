package main

import (
	"context"
	"fmt"
	mis "github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testAuthVal = 0

func testAuthHandler(w http.ResponseWriter, r *http.Request) {
	testAuthVal++
}

type testAuthVerifier struct{}

func (t testAuthVerifier) verify(ctx context.Context, idToken string) (string, error) {
	if idToken == "TestAuthPass" {
		return "user1369", nil
	}
	return "", fmt.Errorf("invalid authentication token")
}

func TestUidInContext(t *testing.T) {
	is := mis.New(t)
	srv, err := newServer()
	is.NoErr(err)

	uid := "user1369"
	modCtx := srv.uidInContext(context.Background(), uid)

	gotUid, err := srv.uidFromContext(modCtx)
	is.NoErr(err)
	is.Equal(gotUid, uid)

	_, err = srv.uidFromContext(context.Background())
	is.True(err != nil)
}

func TestAuth(t *testing.T) {
	is := mis.New(t)

	handlerOpt := func(s *server) {
		s.handler = s.handleAuthed(testAuthHandler)
	}
	srv, err := newServer(handlerOpt, withVerifier(testAuthVerifier{}))
	is.NoErr(err)

	tests := []struct {
		token string
		val   int
		code  int
	}{
		{"Bearer TestAuthPass", 1, http.StatusOK},
		{"Bearer ab76c0239c203c020", 1, http.StatusUnauthorized},
		{"Bearerab76c0239c203c020", 1, http.StatusUnauthorized},
		{" Bearer TestAuthPass", 1, http.StatusUnauthorized},
		{"Bearer		ab76c0239c203c020", 1, http.StatusUnauthorized},
		{"Bearer  TestAuthPass", 2, http.StatusOK},
		{"", 2, http.StatusUnauthorized},
	}

	for _, test := range tests {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", test.token)
		srv.ServeHTTP(rr, req)

		is.Equal(rr.Code, test.code)
		is.Equal(testAuthVal, test.val)
	}
}
