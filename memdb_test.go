package main

import (
	mis "github.com/matryer/is"
	"testing"
	"time"
)

func TestInMemDb(t *testing.T) {
	is := mis.New(t)

	m := newInMemDB()

	_, err := m.Get("noUid")
	if err == nil {
		t.Errorf("Expected error")
	}

	uid := "someUid"

	now := time.Now().UnixNano()
	err = m.Put(uid, Clip{ID: "1001", Timestamp: now})
	is.NoErr(err)

	use := time.Now().UnixNano()
	err = m.Put(uid, Clip{ID: "1002", Timestamp: use})
	is.NoErr(err)

	now =time.Now().UnixNano()
	err = m.Put(uid, Clip{ID: "1003", Timestamp: now})
	is.NoErr(err)

	clips, err := m.Get(uid)
	is.NoErr(err)
	is.Equal(len(clips), 3)
	is.Equal(clips[1].ID, "1002")

	clips, err = m.GetSince(uid, use)
	is.NoErr(err)
	is.Equal(len(clips), 2)
	is.Equal(clips[1].ID, "1003")
}
