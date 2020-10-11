package main

import (
	"github.com/matryer/is"
	"testing"
	"time"
)

func DoTestForDBImpl(t *testing.T, dbImpl DB) {
	is := is.New(t)

	testClips := getSampleTestClips()
	testUser := "user1369"

	// Insert all test clips into db implementation
	for _, testClip := range testClips {
		err := dbImpl.InsertUserClip(testUser, testClip)
		is.NoErr(err)
	}

	// Test Getting from db implementation
	getUserClipsTestCases := []struct {
		description   string
		uid           string
		shouldErr     bool
		expectedCount int
	}{
		{
			"GetUserClipsGoodUID",
			testUser,
			false,
			4,
		},
		{
			"GetUserClipsBadUID",
			"user404",
			true,
			0,
		},
	}

	for _, tc := range getUserClipsTestCases {
		t.Run(tc.description, func(t *testing.T) {
			gotClips, err := dbImpl.GetUserClips(tc.uid)
			if !tc.shouldErr {
				is.NoErr(err)
			} else {
				is.True(err != nil)
			}
			is.Equal(tc.expectedCount, len(gotClips))
		})
	}

	// Test Getting since timestamp from db implementation
	getClipsSinceTestCases := []struct {
		description   string
		uid           string
		timestamp     int64
		expectedCount int
		shouldErr     bool
	}{
		{
			"GetClipsSinceGoodUID",
			testUser,
			testClips[1].Timestamp,
			len(testClips[1:]),
			false,
		},
		{
			"GetClipsSinceBadUID",
			"user404",
			time.Now().UnixNano(),
			0,
			true,
		},
	}

	for _, tc := range getClipsSinceTestCases {
		t.Run(tc.description, func(t *testing.T) {
			gotClips, err := dbImpl.GetUserClipsSince(tc.uid, tc.timestamp)
			if !tc.shouldErr {
				is.NoErr(err)
			} else {
				is.True(err != nil)
			}
			is.Equal(tc.expectedCount, len(gotClips))
		})
	}
}
