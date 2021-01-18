package go_rate

import "sync"

type ArrayMutex struct {
	data      [65536]uint32
	mu        sync.Mutex
	timestamp uint32
}

func NewArrayMutex() *ArrayMutex {
	return &ArrayMutex{}
}

func (r *ArrayMutex) Allow(id uint16, limit, now uint32) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if now > r.timestamp {
		r.timestamp = now

		// clear
		r.data = [65536]uint32{}

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
