package main

type DB interface {
	GetUserClips(userID string) ([]Clip, error)
	GetUserClipsSince(userId string, timestamp int64) ([]Clip, error)
	InsertUserClip(userID string, clip Clip) error
}