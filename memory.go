package goconfig

import (
	"github.com/mekramy/gocast"
)

type memoryDriver struct {
	data map[string]any
}

func (driver *memoryDriver) Load() error {
	if driver.data == nil {
		driver.data = make(map[string]any)
	}
	return nil
}

func (driver *memoryDriver) Set(key string, value any) {
	driver.data[key] = value
}

func (driver *memoryDriver) Get(key string) any {
	if v, ok := driver.data[key]; ok {
		return v
	}

	return nil
}

func (driver *memoryDriver) Exists(key string) bool {
	if _, ok := driver.data[key]; ok {
		return true
	}

	return false
}

func (driver *memoryDriver) Cast(key string) gocast.Caster {
	return gocast.NewCaster(driver.Get(key))
}
