package main

import (
	"context"
	mis "github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testClipsDB struct {}

func (t *testClipsDB) Get(string) ([]Clip, error) {
	return []Clip{}, nil
}

func (t *testClipsDB) GetSince(string, int64) ([]Clip, error) {
	return []Clip{}, nil
}

func (t *testClipsDB) Put(string, Clip) error {
	panic("implement me")
}

func TestClipsGet(t *testing.T) {
	is := mis.New(t)
	srv, err := newServer(withDB(&testClipsDB{}))
	is.NoErr(err)

	req := httptest.NewRequest(http.MethodGet, "/clips", nil)
	badSinceReq := httptest.NewRequest(http.MethodGet, "/clips?since=user136", nil)
	goodSinceReq := httptest.NewRequest(http.MethodGet, "/clips?since=1000000", nil)
	ctxWithUid := context.WithValue(context.Background(), uidKey{}, "user13690")

	testCases := []struct {
		rq   *http.Request
		code int
	}{
		{req, http.StatusInternalServerError},
		{badSinceReq, http.StatusInternalServerError},
		{goodSinceReq, http.StatusInternalServerError},
		{req.WithContext(ctxWithUid), http.StatusOK},
		{badSinceReq.WithContext(ctxWithUid), http.StatusBadRequest},
		{goodSinceReq.WithContext(ctxWithUid), http.StatusOK},
	}

	for _, tc := range testCases {
		rr := httptest.NewRecorder()
		srv.handleClipsGet().ServeHTTP(rr, tc.rq)

		is.Equal(rr.Code, tc.code)
	}
}
