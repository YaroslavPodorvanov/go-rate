package go_rate

import "sync"

type MapMutex struct {
	data      map[uint16]uint32
	mu        sync.Mutex
	timestamp uint32
}

func NewMapMutex() *MapMutex {
	return &MapMutex{
		data: make(map[uint16]uint32),
	}
}

func (r *MapMutex) Allow(id uint16, limit, now uint32) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if now > r.timestamp {
		r.timestamp = now

		// clear
		for key := range r.data {
			delete(r.data, key)
		}

		r.data[id] = 1

		return true
	}

	var value = r.data[id]

	if value >= limit {
		return false
	}

	r.data[id] = value + 1

	return true
}
