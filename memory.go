package goconfig

import (
	"sync"

	"github.com/mekramy/gocast"
)

type memoryDriver struct {
	data  map[string]any
	mutex sync.RWMutex
}

func (m *memoryDriver) Load() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.data == nil {
		m.data = make(map[string]any)
	}
	return nil
}

func (m *memoryDriver) Set(key string, value any) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data[key] = value
}

func (m *memoryDriver) Get(key string) any {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if v, ok := m.data[key]; ok {
		return v
	}

	return nil
}

func (m *memoryDriver) Exists(key string) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if _, ok := m.data[key]; ok {
		return true
	}

	return false
}

func (m *memoryDriver) Cast(key string) gocast.Caster {
	return gocast.NewCaster(m.Get(key))
}
