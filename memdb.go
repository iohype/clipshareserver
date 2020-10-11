package main

import (
	"fmt"
	"sync"
)

type inMemDb struct {
	mu sync.RWMutex
	// map a uid to their clips
	data map[string][]Clip
}

func newInMemDB() *inMemDb {
	return &inMemDb{
		data: make(map[string][]Clip),
	}
}

func (m *inMemDb) Get(uid string) ([]Clip, error) {
	m.mu.RLock()
	clips, ok := m.data[uid]
	m.mu.RUnlock()
	if !ok {
		return []Clip{}, fmt.Errorf("uid does not exist")
	}
	return clips, nil
}

func (m *inMemDb) GetSince(userID string, timestamp int64) ([]Clip, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	clipsForUser, ok := m.data[userID]
	if !ok {
		return []Clip{}, fmt.Errorf("uid does not exist")
	}

	var clipsSinceTimestamp = make([]Clip, 0)
	for _, clip := range clipsForUser {
		if clip.Timestamp >= timestamp {
			clipsSinceTimestamp = append(clipsSinceTimestamp, clip)
		}
	}
	return clipsSinceTimestamp, nil
}

func (m *inMemDb) Put(uid string, clip Clip) error {
	if uid == "" {
		return fmt.Errorf("invalid uid")
	}
	if clip.Data == "" || clip.Timestamp == 0 || clip.ID == "" {
		return fmt.Errorf("clip is empty")
	}
	if clip.Timestamp < 0 {
		return fmt.Errorf("clip has negative time")
	}
	m.mu.Lock()
	m.data[uid] = append(m.data[uid], clip)
	m.mu.Unlock()
	return nil
}