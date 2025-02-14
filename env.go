package goconfig

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/mekramy/gocast"
)

type envDriver struct {
	files []string
	data  map[string]any
	mutex sync.RWMutex
}

func (e *envDriver) Load() error {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if e.data == nil {
		e.data = make(map[string]any)
	}

	return godotenv.Overload(e.files...)
}

func (e *envDriver) Set(key string, value any) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.data[key] = value
}

func (e *envDriver) Get(key string) any {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	if v, ok := e.data[key]; ok {
		return v
	}

	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return nil
}

func (e *envDriver) Exists(key string) bool {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	if _, ok := e.data[key]; ok {
		return true
	}

	if _, ok := os.LookupEnv(key); ok {
		return true
	}

	return false
}

func (e *envDriver) Cast(key string) gocast.Caster {
	return gocast.NewCaster(e.Get(key))
}
