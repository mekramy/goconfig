package goconfig

import (
	"sync"

	"github.com/mekramy/gocast"
)

type memoryDriver struct {
	data  map[string]any
	mutex sync.RWMutex
}

func (driver *memoryDriver) Load() error {
	driver.mutex.Lock()
	defer driver.mutex.Unlock()

	if driver.data == nil {
		driver.data = make(map[string]any)
	}
	return nil
}

func (driver *memoryDriver) Set(key string, value any) {
	driver.mutex.Lock()
	defer driver.mutex.Unlock()

	driver.data[key] = value
}

func (driver *memoryDriver) Get(key string) any {
	driver.mutex.RLock()
	defer driver.mutex.RUnlock()

	if v, ok := driver.data[key]; ok {
		return v
	}

	return nil
}

func (driver *memoryDriver) Exists(key string) bool {
	driver.mutex.RLock()
	defer driver.mutex.RUnlock()

	if _, ok := driver.data[key]; ok {
		return true
	}

	return false
}

func (driver *memoryDriver) Cast(key string) gocast.Caster {
	return gocast.NewCaster(driver.Get(key))
}
