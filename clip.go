package main

// A clipboard item
type Clip struct {
	ID        string `json:"id"`
	Timestamp int64 `json:"timestamp"`
	Data      string `json:"data"`
}
