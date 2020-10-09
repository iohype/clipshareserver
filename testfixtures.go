package main

import (
	"time"
)

func getSampleTestClips() []Clip {
	testTime := time.Now()
	return []Clip{
		{"1001", testTime.UnixNano(), "test data content"},
		{"1002", testTime.Add(2 * time.Microsecond).UnixNano(), "test data"},
		{"1004", testTime.Add(1 * time.Hour).UnixNano(), "test"},
		{"1005", testTime.Add(3 * time.Hour).UnixNano(), "data"},
	}
}

func getTestMemDB() (*inMemDb, error) {
	m := newInMemDB()
	testClips := getSampleTestClips()
	testUser := "user1369"
	for _, testClip := range testClips {
		err := m.InsertUserClip(testUser, testClip)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}