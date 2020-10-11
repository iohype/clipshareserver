package main

import (
	"net/http"
	"strings"
)

func (s *server) requireAuthed(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idToken := tokenFromHeader(r)

		userID, err := s.verifier.verify(idToken)
		if err != nil {
			s.Error(w, err, http.StatusUnauthorized)
			return
		}

		ctxWithUserID := putUserIDInContext(r.Context(), userID)
		next(w, r.WithContext(ctxWithUserID))
	}
}

func tokenFromHeader(r *http.Request) string {
	idToken := r.Header.Get("Authorization")
	idToken = strings.TrimPrefix(idToken, "Bearer")
	idToken = strings.TrimSpace(idToken)
	return idToken
}
