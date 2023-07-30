package lib

import (
	"sync"

	sets "github.com/daichi-m/go18ds/sets/hashset"
)

type Set[A comparable] struct {
	m  *sets.Set[A]
	mu *sync.RWMutex
}

func NewSet[A comparable]() *Set[A] {
	return &Set[A]{
		m:  sets.New[A](),
		mu: &sync.RWMutex{},
	}
}

func (m *Set[A]) Get(a A) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m.Contains(a)
}

func (m *Set[A]) Keys() []A {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m.Values()
}

func (m *Set[A]) Remove(a A) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m.Remove(a)
}

func (m *Set[A]) Set(a A) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m.Add(a)
}
