package main

import (
	"net/http"
	"strconv"
)

func (s *server) handleClipsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		since := r.URL.Query().Get("since")

		// Get userID from context
		userID, err := s.uidFromContext(r.Context())
		if err != nil {
			s.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		var clips []Clip
		if since != "" {
			s.logger.Printf("Getting clips for uid %s since %s\n", userID, since)
			timestamp, err := strconv.ParseInt(since, 10, 64)
			if err != nil {
				s.Error(w, r, err, http.StatusBadRequest)
			}
			clips, err = s.db.GetSince(userID, timestamp)
		} else {
			s.logger.Printf("Getting all clips for uid %s", userID)
			clips, err = s.db.Get(userID)
		}
		if err != nil {
			s.Error(w, r, err, http.StatusInternalServerError)
			return
		}

		s.JSON(w, r, clips, http.StatusOK)
	}
}
