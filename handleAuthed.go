package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// Key used to retrieve uid from context
type uidKey struct{}

//handleAuthed is a middleware that allows the request only if user is logged in
func (s *server) handleAuthed(vfy verifier, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idToken := tokenFromHeader(r)
		// Verify token, verifier should handle caching itself
		uid, err := vfy.verify(r.Context(), idToken)
		if err != nil {
			s.Error(w, r, err, http.StatusUnauthorized)
			return
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
