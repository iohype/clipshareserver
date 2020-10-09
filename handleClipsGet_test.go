package main

import (
	"context"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleClipsGet(t *testing.T) {
	is := is.New(t)
	m, err := getTestMemDB()
	is.NoErr(err)
	srv, err := newServer(withDB(m))
	is.NoErr(err)

	req := httptest.NewRequest(http.MethodGet, "/clips", nil)
	badSinceReq := httptest.NewRequest(http.MethodGet, "/clips?since=abcdefg", nil)
	goodSinceReq := httptest.NewRequest(http.MethodGet, "/clips?since=1000000", nil)
	ctx := context.Background()

	userWithData := "user1369"
	userWithNoData := "userEmpty"

	testCases := []struct {
		description  string
		req          *http.Request
		expectedCode int
	}{
		{
			"GetClipsUIDHasDataTest",
			req.WithContext(putUserIDInContext(ctx, userWithData)),
			http.StatusOK,
		},
		{
			"GetClipsUIDHasNoDataTest",
			req.WithContext(putUserIDInContext(ctx, userWithNoData)),
			http.StatusOK,
		},
		{
			"GetClipsBadSinceQueryTest",
			badSinceReq.WithContext(putUserIDInContext(ctx, userWithData)),
			http.StatusBadRequest,
		},
		{
			"GetClipsGoodSinceUIDHasDataTest",
			goodSinceReq.WithContext(putUserIDInContext(ctx, userWithData)),
			http.StatusOK,
		},
		{
			"GetClipsGoodSinceUIDHasNoDataTest",
			goodSinceReq.WithContext(putUserIDInContext(ctx, userWithNoData)),
			http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			rr := httptest.NewRecorder()
			srv.handleClipsGet().ServeHTTP(rr, tc.req)
			is.Equal(tc.expectedCode, rr.Code)
		})
	}
}
