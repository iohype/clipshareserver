package main

import (
	"github.com/matryer/is"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

type testOptVerifier struct {}

func (t *testOptVerifier) verify(string) (string, error) {
	panic("noop")
}

func testOptHandler(http.ResponseWriter, *http.Request) {}

func TestServerOptConfig(t *testing.T) {
	is := is.New(t)

	tVerifier := &testOptVerifier{}
	tLogger := log.New(ioutil.Discard, "", 0)
	tAddr := ":5050"
	tHandler := http.HandlerFunc(testOptHandler)
	tDB := newInMemDB()

	srv, err := newServer(
		withVerifier(tVerifier),
		withLogger(tLogger),
		withAddr(tAddr),
		withHandler(tHandler),
		withDB(tDB))

	is.NoErr(err)

	is.Equal(srv.verifier, tVerifier)
	is.Equal(srv.logger, tLogger)
	is.Equal(srv.Addr, tAddr)
	is.Equal(srv.handler, tHandler)
	is.True(srv.db == tDB)
}