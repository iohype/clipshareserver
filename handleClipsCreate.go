package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (s *server) handleClipsCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse and decode the request body into a new `Clip` instance
		var clip Clip

		err := json.NewDecoder(r.Body).Decode(&clip)
		if err != nil {
			// If there is something wrong with the request body, return a 400 status
			s.Error(w, err, http.StatusBadRequest)
			return
		}

		userID, err := getUserIDFromContext(r.Context())
		if err != nil {
			s.Error(w, err, http.StatusInternalServerError)
			return
		}

		clip.Timestamp = time.Now().UnixNano()

		if err = s.db.InsertUserClip(userID, clip); err != nil {
			// If there is any issue with inserting into the database, return a 500 error
			s.Error(w, err, http.StatusInternalServerError)
			return
		}
		s.JSON(w, clip, http.StatusCreated)
	}
}
