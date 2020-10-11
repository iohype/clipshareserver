package main

import (
	"log"
	"net"
	"os"
)

const defaultRunAddr = ":8080"

func main() {
	err := runApp()
	if err != nil {
		log.Fatal(err)
	}
}

func runApp() error {
	srv, err := getRunServer()
	if err != nil {
		return err
	}
	return start(srv)
}

func getRunServer() (*server, error) {
	runAddr := getRunAddr()
	return newServer(withAddr(runAddr))
}

func getRunAddr() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return defaultRunAddr
	}
	return addrFromPort(port)
}

func addrFromPort(port string) string {
	return net.JoinHostPort("", port)
}
