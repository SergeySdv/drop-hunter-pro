package lib

import (
	"sync"

	maps "github.com/daichi-m/go18ds/maps/hashmap"
)

type Map[A comparable, B comparable] struct {
	m  *maps.Map[A, B]
	mu *sync.RWMutex
}

func NewMap[A comparable, B comparable]() *Map[A, B] {

	return &Map[A, B]{
		m:  maps.New[A, B](),
		mu: &sync.RWMutex{},
	}
}

func (m *Map[A, B]) Get(a A) (B, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m.Get(a)
}

func (m *Map[A, B]) Keys() []A {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m.Keys()
}

func (m *Map[A, B]) Remove(a A) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m.Remove(a)
}

func (m *Map[A, B]) Set(a A, b B) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m.Put(a, b)
}
