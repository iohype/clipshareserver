package main

import (
	"bytes"
	"errors"
	mis "github.com/matryer/is"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_Error(t *testing.T) {
	is := mis.New(t)
	srv, err := newServer()
	is.NoErr(err)

	testCases := []struct {
		description string
		err         error
		code        int
		expected    string
	}{
		{
			"GoodParamsTest",
			errors.New("something bad happened"),
			http.StatusNotImplemented,
			`{"message":"something bad happened"}`,
		},
		{
			"NilErrorTest",
			nil,
			http.StatusNotImplemented,
			`{"message":null}`,
		},
		{
			"BizarreStatusAndErrTest",
			nil,
			1000,
			`{"message":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			rr := httptest.NewRecorder()
			srv.Error(rr, tc.err, tc.code)
			gotJSON, err := ioutil.ReadAll(rr.Body)
			is.NoErr(err)

			gotJSON = bytes.TrimSpace(gotJSON)
			is.Equal(tc.code, rr.Code)
			is.Equal(tc.expected, string(gotJSON))
		})
	}

}

func TestServer_JSON(t *testing.T) {
	is := mis.New(t)
	srv, err := newServer()
	is.NoErr(err)

	testCases := []struct{
		description string
		data interface{}
		code int
		expected string
	}{
		{
			"GoodParamStringTest",
			"TestData",
			http.StatusNotImplemented,
			`"TestData"`,
		},
		{
			"GoodParamStructTest",
			&struct {
				Value int64	`json:"value"`
			}{
				42000000000000000,
			},
			http.StatusNotImplemented,
			`{"value":42000000000000000}`,
		},
		{
			"NilParamTest",
			nil,
			http.StatusNotImplemented,
			`null`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			rr := httptest.NewRecorder()
			srv.JSON(rr, tc.data, tc.code)
			gotJSON, err := ioutil.ReadAll(rr.Body)
			is.NoErr(err)

			gotJSON = bytes.TrimSpace(gotJSON)
			is.Equal("application/json", rr.Header().Get("Content-Type"))
			is.Equal(tc.code, rr.Code)
			is.Equal(tc.expected, string(gotJSON))
		})
	}
}
