package main

import (
	"context"
	"fmt"
)

type uidKey struct{}

func putUserIDInContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, uidKey{}, userID)
}

func getUserIDFromContext(ctx context.Context) (userID string, err error) {
	val, ok := ctx.Value(uidKey{}).(string)
	if !ok {
		return "", fmt.Errorf("invalid or no UserID in context")
	}
	return val, nil
}
