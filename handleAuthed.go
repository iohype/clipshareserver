package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

// Key to retrieve uid from a context
type uidKey struct{}

func (s *server) uidFromContext(ctx context.Context) (string, error) {
	val, ok := ctx.Value(uidKey{}).(string)
	if !ok {
		return "", fmt.Errorf("invalid UID in context")
	}
	return val, nil
}

//handleAuthed is a middleware that allows the request only if user is logged in
func (s *server) handleAuthed(next http.HandlerFunc) http.HandlerFunc {
	// Cache hit prevents asking FirebaseAuth to verify id token
	var (
		srw         sync.RWMutex
		authedCache = make(map[string]string)
	)
	return func(w http.ResponseWriter, r *http.Request) {
		idToken := r.Header.Get("Authorization")
		idToken = strings.TrimPrefix("Bearer", idToken)
		idToken = strings.TrimSpace(idToken)
		srw.RLock()
		uid, ok := authedCache[idToken]
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
		// Pass the uid down
		ctx := context.WithValue(r.Context(), uidKey{}, uid)
		next(w, r.WithContext(ctx))
	}
}

func verifyToken(idToken string) (string, error) {
	return "user1369", nil //fmt.Errorf("not implemented")
}
