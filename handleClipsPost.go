package main

import (
	"net/http"
	"encoding/json"
	"time"
)

func (s *server) handleClipsPost(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Clip` instance
	var clip Clip 

	err := json.NewDecoder(r.Body).Decode(&clip)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		s.Error(w, r, err, http.StatusBadRequest)
		return
	}

	uid, err := s.uidFromContext(r.Context())
	if err != nil {
		s.Error(w, r, err, http.StatusInternalServerError)
		return
	}
	
	clip.Timestamp = time.Now().UnixNano() 

	if err = s.db.Put(uid, clip); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		s.Error(w, r, err, http.StatusInternalServerError)
		return
	}
	s.JSON(w, r, clip, http.StatusCreated) 
}