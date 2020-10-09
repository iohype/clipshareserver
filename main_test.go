package main

import (
	"github.com/matryer/is"
	"os"
	"testing"
)

var testPort = "3333"

func doEnvSetup(is *is.I) func() {
	portKey := "PORT"
	err := os.Setenv(portKey, testPort)
	is.NoErr(err)

	return func() {
		err := os.Unsetenv(portKey)
		is.NoErr(err)
	}
}

func TestRunAddrFromEnvPort(t *testing.T) {
	is := is.New(t)

	cleanup := doEnvSetup(is)
	defer cleanup()

	gotAddr := getRunAddr()
	is.Equal(gotAddr, ":"+testPort)
}

func TestDefaultRunAddr(t *testing.T) {
	is := is.New(t)

	gotAddr := getRunAddr()
	is.Equal(gotAddr, defaultRunAddr)
}
