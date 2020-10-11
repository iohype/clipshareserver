package main

import (
	"net/http"
	"strconv"
)

func (s *server) handleClipsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := getUserIDFromContext(r.Context())
		if err != nil {
			s.Error(w, err, http.StatusInternalServerError)
			return
		}

		sinceParam := r.URL.Query().Get("since")
		if sinceParam != "" {
			s.handleGetUserClipsSince(w, userID, sinceParam)
		} else {
			s.handleGetAllUserClips(w, userID)
		}
	}
}

func (s *server) handleGetUserClipsSince(w http.ResponseWriter, userID string, sinceParam string) {
	s.logger.Printf("Getting clips for userID %s since %s\n", userID, sinceParam)

	timestamp, err := paramToTimestamp(sinceParam)
	if err != nil {
		s.Error(w, err, http.StatusBadRequest)
		return
	}

	clips, err := s.db.GetUserClipsSince(userID, timestamp)
	if err != nil {
		s.logger.Println(err)
		clips = []Clip{}
	}

	s.JSON(w, clips, http.StatusOK)
}

func (s *server) handleGetAllUserClips(w http.ResponseWriter, userID string) {
	s.logger.Printf("Getting all clips for userID %s", userID)

	clips, err := s.db.GetUserClips(userID)
	if err != nil {
		s.logger.Println(err)
		clips = []Clip{}
	}

	s.JSON(w, clips, http.StatusOK)
}

func paramToTimestamp(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

