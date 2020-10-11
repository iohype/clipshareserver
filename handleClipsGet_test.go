package main

import (
	"context"
	mis "github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testDataClipsGet = map[string][]Clip{
	"user1369": {
		{
			"1001",
			currentTimeUnixNano(),
			"test data content",
		},
		{
			"1002",
			currentTimeUnixNano(),
			"test data content",
		},
		{
			"1005",
			currentTimeUnixNano(),
			"test data other content",
		},
	},
	"user42": {},
}

func TestClipsGet(t *testing.T) {
	is := mis.New(t)
	m := newInMemDB()
	m.data = testDataClipsGet
	srv, err := newServer(withDB(m))
	is.NoErr(err)

	req := httptest.NewRequest(http.MethodGet, "/clips", nil)
	badSinceReq := httptest.NewRequest(http.MethodGet, "/clips?since=abcdefg", nil)
	goodSinceReq := httptest.NewRequest(http.MethodGet, "/clips?since=1000000", nil)
	ctxWithUid := context.WithValue(context.Background(), uidKey{}, "user1369")

	testCases := []struct {
		description string
		rq       *http.Request
		code     int
	}{
		{
			"GoodRequestNoContextUidTest",
			req,
			http.StatusInternalServerError,
		},
		{
			"BadSinceRequestNoContextUidTest",
			badSinceReq,
			http.StatusInternalServerError,
		},
		{
			"GoodSinceRequestNoContextUidTest",
			goodSinceReq,
			http.StatusInternalServerError,
		},
		{
			"GoodRequestWithContextUidTest",
			req.WithContext(ctxWithUid),
			http.StatusOK,
		},
		{
			"BadSinceRequestWithContextUidTest",
			badSinceReq.WithContext(ctxWithUid),
			http.StatusBadRequest,
		},
		{
			"GoodSinceRequestWithContextUidTest",
			goodSinceReq.WithContext(ctxWithUid),
			http.StatusOK,
		},
	}

	for _, tc := range testCases {
		rr := httptest.NewRecorder()
		srv.handleClipsGet().ServeHTTP(rr, tc.rq)

		is.Equal(rr.Code, tc.code)
	}
}
