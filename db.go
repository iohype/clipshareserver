package main

//DB defines methods for a backing database
type DB interface {
	// get all user clips
	Get(userID string) ([]Clip, error)
	// get all user clips since timestamp
	GetSince(userId string, timestamp int64) ([]Clip, error)
	// save a clip
	Put(userID string, clip Clip) error
}
