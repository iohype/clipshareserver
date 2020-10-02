package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *server) Error(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	v := &struct {
		Error string `json:"error"`
	}{
		err.Error(),
	}
	s.JSON(w, r, v, statusCode)
}

//JSON writes data out in JSON format
func (s *server) JSON(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}