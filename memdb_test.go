package main

import "testing"

func TestMemDB(t *testing.T) {
	DoTestForDBImpl(t, newInMemDB())
}
