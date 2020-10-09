package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleClipsCreate(t *testing.T) {
	is := is.New(t)
	m := newInMemDB()
	srv, err := newServer(withDB(m))
	is.NoErr(err)

	testUser := "user1369"

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&Clip{Data: "Test Clip Content"})
	is.NoErr(err)

	req := httptest.NewRequest(http.MethodPost, "/clips", &b)
	reqWithUid := req.WithContext(putUserIDInContext(context.Background(), testUser))
	rr := httptest.NewRecorder()
	srv.handleClipsCreate().ServeHTTP(rr, reqWithUid)

	is.Equal(rr.Code, http.StatusCreated)
	userClips, err := m.GetUserClips(testUser)
	is.NoErr(err)
	is.Equal(len(userClips), 1)
	is.Equal(userClips[0].Data, "Test Clip Content")
}
