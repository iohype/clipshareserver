package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

// Key used to retrieve uid from context
type uidKey struct{}

//handleAuthed is a middleware that allows the request only if user is logged in
func (s *server) handleAuthed(next http.HandlerFunc) http.HandlerFunc {
	// Cache hit prevents asking FirebaseAuth to verify id token
	var (
		mu          sync.RWMutex
		authedCache = make(map[string]string)
	)
	return func(w http.ResponseWriter, r *http.Request) {
		idToken := tokenFromHeader(r)
		// Read from cache
		mu.RLock()
		uid, ok := authedCache[idToken]
		mu.RUnlock()
		if !ok {
			// token not in cache, try FirebaseAuth
			// err is declared this way here to avoid creating a new uid variable
			//in this if block due to the use of the := shortcut
			var err error
			uid, err = verifyToken(idToken)
			if err != nil {
				s.Error(w, r, err, http.StatusUnauthorized)
				return
			}
			// Save to cache
			mu.Lock()
			authedCache[idToken] = uid
			mu.Unlock()
		}
		// Pass the uid down to handlers
		ctx := s.uidInContext(r.Context(), uid)
		next(w, r.WithContext(ctx))
	}
}

//uidInContext returns a context cloned from the baseCtx context
//but with uidKey{} set to value uid
func (s *server) uidInContext(baseCtx context.Context, uid string) context.Context {
	return context.WithValue(baseCtx, uidKey{}, uid)
}

//uidFromContext returns the uid from ctx context
func (s *server) uidFromContext(ctx context.Context) (string, error) {
	val, ok := ctx.Value(uidKey{}).(string)
	if !ok {
		return "", fmt.Errorf("invalid UID in context")
	}
	return val, nil
}

func tokenFromHeader(r *http.Request) string {
	idToken := r.Header.Get("Authorization")
	idToken = strings.TrimPrefix(idToken, "Bearer")
	idToken = strings.TrimSpace(idToken)
	return idToken
}

func verifyToken(idToken string) (string, error) {
	return "user1369", nil //fmt.Errorf("not implemented")
}
