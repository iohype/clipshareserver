package main

import (
	"errors"
	mis "github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_Error(t *testing.T) {
	is := mis.New(t)

	srv := &server{}
	rr := httptest.NewRecorder()
	testErr := errors.New("something bad happened")
	srv.Error(rr, nil, testErr, http.StatusNotImplemented)

	is.Equal(rr.Code, http.StatusNotImplemented)
}

func TestServer_JSON(t *testing.T) {
	is := mis.New(t)

	srv := &server{}
	rr := httptest.NewRecorder()
	testData := "Data"
	srv.JSON(rr, nil, testData, http.StatusNotImplemented)

	is.Equal(rr.Code, http.StatusNotImplemented)
	is.Equal("application/json", rr.Header().Get("Content-Type"))
}