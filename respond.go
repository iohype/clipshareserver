package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type errResponse struct {
	Message interface{} `json:"message"`
}

func (s *server) Error(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	s.logger.Println(err)
	s.JSON(w, r, toErrorResponse(err), statusCode)
}

//JSON writes data out in JSON format
func (s *server) JSON(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}

func toErrorResponse(err error) errResponse {
	if err != nil {
		return errResponse{err.Error()}
	}
	return errResponse{nil}
}
