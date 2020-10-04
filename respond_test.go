package main

import (
	"bytes"
	"encoding/json"
	"errors"
	mis "github.com/matryer/is"
	"io/ioutil"
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

	data, err := json.Marshal(testData)
	is.NoErr(err)
	recvData, err := ioutil.ReadAll(rr.Body)
	is.NoErr(err)
	is.Equal(data, bytes.TrimSpace(recvData))
}
