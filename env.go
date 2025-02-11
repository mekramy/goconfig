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

func (driver *envDriver) Load() error {
	driver.mutex.Lock()
	defer driver.mutex.Unlock()

	if driver.data == nil {
		driver.data = make(map[string]any)
	}

	return godotenv.Overload(driver.files...)
}

func (driver *envDriver) Set(key string, value any) {
	driver.mutex.Lock()
	defer driver.mutex.Unlock()

	driver.data[key] = value
}

func (driver *envDriver) Get(key string) any {
	driver.mutex.RLock()
	defer driver.mutex.RUnlock()

	if v, ok := driver.data[key]; ok {
		return v
	}

	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return nil
}

func (driver *envDriver) Exists(key string) bool {
	driver.mutex.RLock()
	defer driver.mutex.RUnlock()

	if _, ok := driver.data[key]; ok {
		return true
	}

	if _, ok := os.LookupEnv(key); ok {
		return true
	}

	return false
}

func (driver *envDriver) Cast(key string) gocast.Caster {
	return gocast.NewCaster(driver.Get(key))
}
