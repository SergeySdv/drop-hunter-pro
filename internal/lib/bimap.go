package lib

import (
	"sync"

	maps "github.com/daichi-m/go18ds/maps/hashbidimap"
)

type BiMap[A comparable, B comparable] struct {
	m  *maps.Map[A, B]
	mu *sync.RWMutex
}

func NewBiMap[A comparable, B comparable]() *BiMap[A, B] {

	return &BiMap[A, B]{
		m:  maps.New[A, B](),
		mu: &sync.RWMutex{},
	}
}

func (m *BiMap[A, B]) Get(a A) (B, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m.Get(a)
}

func (m *BiMap[A, B]) Keys() []A {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m.Keys()
}

func (m *BiMap[A, B]) GetKey(b B) (A, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m.GetKey(b)
}

func (m *BiMap[A, B]) Remove(a A) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m.Remove(a)
}

func (m *BiMap[A, B]) Set(a A, b B) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m.Put(a, b)
}
