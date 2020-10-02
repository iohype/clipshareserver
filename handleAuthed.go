package main

import (
	"net/http"
	"strings"
	"sync"
)

//handleAuthed is a middleware that allows the request only if user is logged in
func (s *server) handleAuthed(next http.HandlerFunc) http.HandlerFunc {
	// Cache hit prevents asking FirebaseAuth to verify id token
	var (
		srw sync.RWMutex
		authedCache = make(map[string]string)
	)
	return func(w http.ResponseWriter, r *http.Request) {
		idToken := r.Header.Get("Authorization")
		idToken = strings.TrimPrefix("Bearer", idToken)
		idToken = strings.TrimSpace(idToken)
		srw.RLock()
		_, ok := authedCache[idToken]
		srw.RUnlock()
		if !ok {
			// idToken not in cache, try firebase auth
			uid, err := verifyToken(idToken)
			if err != nil {
				s.Error(w, r, err, http.StatusUnauthorized)
				return
			}
			srw.Lock()
			authedCache[idToken] = uid
			srw.Unlock()
		}
		next(w, r)
	}
}

func verifyToken(idToken string) (string, error) {
	return "", nil
}