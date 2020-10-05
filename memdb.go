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
		return clips, fmt.Errorf("uid does not exist")
	}
	return clips, nil
}

func (m *inMemDb) GetSince(uid string, since int64) ([]Clip, error) {
	var clips []Clip
	uidClips, err := m.Get(uid)
	if err != nil {
		return uidClips, err
	}
	for _, clip := range uidClips {
		if clip.Timestamp >= since {
			clips = append(clips, clip)
		}
	}
	return clips, nil
}

func (m *inMemDb) Put(uid string, clip Clip) error {
	m.mu.Lock()
	m.data[uid] = append(m.data[uid], clip)
	m.mu.Unlock()
	return nil
}