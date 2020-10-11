package main

import (
	mis "github.com/matryer/is"
	"testing"
	"time"
)

var testDataMemdb = map[string][]Clip{
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

func currentTimeUnixNano() int64 {
	return time.Now().UnixNano()
}

func TestInMemDb_Put(t *testing.T) {
	is := mis.New(t)
	m := newInMemDB()
	m.data = testDataMemdb

	testCases := []struct {
		description   string
		uid           string
		clip          Clip
		expectedCount int
		shouldErr     bool
	}{
		{
			"GoodUidGoodClipTest",
			"user1369",
			Clip{"1007", currentTimeUnixNano(), "clip content"},
			4,
			false,
		},
		{
			"EmptyUidGoodClipTest",
			"",
			Clip{"0", currentTimeUnixNano(), "clip content"},
			0,
			true,
		},
		{
			"GoodUidEmptyClipTest",
			"user1369",
			Clip{},
			4,
			true,
		},
		{
			"GoodUidBadClipTest",
			"user1369",
			Clip{"1", -5000, "clip content"},
			4,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			err := m.Put(tc.uid, tc.clip)
			if tc.shouldErr {
				is.True(err != nil)
			} else {
				is.NoErr(err)
			}
			is.Equal(tc.expectedCount, len(m.data[tc.uid]))
		})
	}
}

func TestInMemDb_GetSince(t *testing.T) {
	is := mis.New(t)
	m := newInMemDB()
	m.data = testDataMemdb

	testCases := []struct {
		description string
		uid         string
		expected    []Clip
		shouldErr   bool
		since       int64
	}{
		{
			"GoodUidParamTest",
			"user1369",
			testDataMemdb["user1369"][1:],
			false,
			testDataMemdb["user1369"][1].Timestamp,
		},
		{
			"GoodUidTooRecentTimestampTest",
			"user1369",
			[]Clip{},
			false,
			currentTimeUnixNano(),
		},
		{
			"GoodUidParamNoDataTest",
			"user42",
			[]Clip{},
			false,
			currentTimeUnixNano(),
		},
		{
			"NonExistentUidParamTest",
			"user404",
			[]Clip{},
			true,
			currentTimeUnixNano(),
		},
		{
			"EmptyUidParamTest",
			"",
			[]Clip{},
			true,
			currentTimeUnixNano(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			gotClips, err := m.GetSince(tc.uid, tc.since)
			if tc.shouldErr {
				is.True(err != nil)
			} else {
				is.NoErr(err)
			}
			is.Equal(gotClips, tc.expected)
		})
	}
}

func TestInMemDb_Get(t *testing.T) {
	is := mis.New(t)
	m := newInMemDB()
	m.data = testDataMemdb

	testCases := []struct {
		description string
		uid         string
		expected    []Clip
		shouldErr   bool
	}{
		{
			"GoodUidParamTest",
			"user1369",
			testDataMemdb["user1369"],
			false,
		},
		{
			"GoodUidParamNoDataTest",
			"user42",
			[]Clip{},
			false,
		},
		{
			"NonExistentUidParamTest",
			"user404",
			[]Clip{},
			true,
		},
		{
			"EmptyUidParamTest",
			"",
			[]Clip{},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			gotClips, err := m.Get(tc.uid)
			if tc.shouldErr {
				is.True(err != nil)
			} else {
				is.NoErr(err)
			}
			is.Equal(gotClips, tc.expected)
		})
	}
}

