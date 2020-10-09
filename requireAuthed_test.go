package main

import (
	"fmt"
	mis "github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)


func testAuthHandler(w http.ResponseWriter, r *http.Request) {}

type testAuthVerifier struct{}

func (t testAuthVerifier) verify(idToken string) (string, error) {
	if idToken == "TestAuthPass" {
		return "user1369", nil
	}
	return "", fmt.Errorf("invalid authentication token")
}

func TestRequireAuthed(t *testing.T) {
	is := mis.New(t)

	handlerOpt := func(s *server) {
		s.handler = s.requireAuthed(testAuthHandler)
	}
	srv, err := newServer(handlerOpt, withVerifier(testAuthVerifier{}))
	is.NoErr(err)

	tests := []struct {
		token string
		code  int
	}{
		{"Bearer TestAuthPass", http.StatusOK},
		{"Bearer ab76c0239c203c020", http.StatusUnauthorized},
		{"Bearerab76c0239c203c020", http.StatusUnauthorized},
		{" Bearer TestAuthPass", http.StatusUnauthorized},
		{"Bearer		ab76c0239c203c020",http.StatusUnauthorized},
		{"Bearer  TestAuthPass", http.StatusOK},
		{"", http.StatusUnauthorized},
	}

	for _, test := range tests {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", test.token)
		srv.ServeHTTP(rr, req)

		is.Equal(rr.Code, test.code)
	}
}
