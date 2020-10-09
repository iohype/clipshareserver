package main

import (
	"context"
	mis "github.com/matryer/is"
	"testing"
)

func TestUserIDInContext(t *testing.T) {
	is := mis.New(t)

	ctx := context.Background()
	testUid := "user1369"
	testCases := []struct {
		description string
		testCtx     context.Context
		shouldErr   bool
		expected    string
	}{
		{
			"UidInContextTest",
			putUserIDInContext(ctx, testUid),
			false,
			testUid,
		},
		{
			"UidNotInContextTest",
			ctx,
			true,
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			gotUid, err := getUserIDFromContext(tc.testCtx)
			if !tc.shouldErr {
				is.NoErr(err)
			} else {
				is.True(err != nil)
			}
			is.Equal(tc.expected, gotUid)
		})
	}
}
