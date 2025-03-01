package memory

import "sync"

type Memory[T comparable, R any] struct {
	buffer map[T]R
	mu     sync.RWMutex
}

func New[T comparable, R any]() *Memory[T, R] {
	return &Memory[T, R]{
		buffer: make(map[T]R),
	}
}

func (m *Memory[T, R]) Store(key T, value R) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.buffer[key] = value
}

func (m *Memory[T, R]) Get(key T) (R, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, exist := m.buffer[key]

	return value, exist
}
