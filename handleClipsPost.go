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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uid := clip.ID
	clip.Timestamp = time.Now().UnixNano() 

	if err = s.db.Put(uid, clip); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		s.logger.Printf("There was an issue %s\n", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(clip)
}