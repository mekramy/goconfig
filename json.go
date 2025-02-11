package goconfig

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/mekramy/gocast"
	"github.com/tidwall/gjson"
)

type jsonDriver struct {
	raw   string
	files []string
	data  map[string]any
	mutex sync.RWMutex
}

func (driver *jsonDriver) Load() error {
	driver.mutex.Lock()
	defer driver.mutex.Unlock()

	if driver.data == nil {
		driver.data = make(map[string]any)
	}

	// Read json files
	contents := make([]string, 0)
	for _, file := range driver.files {
		bytes, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		content := string(bytes)
		if !gjson.Valid(content) {
			return fmt.Errorf("invalid json in %s", file)
		}

		fileName := filepath.Base(file)
		fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

		if len(driver.files) > 1 {
			contents = append(contents, `"`+fileName+`":`+content)
		} else {
			contents = append(contents, content)
		}
	}

	// Generate big config file
	if len(driver.files) > 1 {
		driver.raw = "{" + strings.Join(contents, ",") + "}"
	} else if len(driver.files) > 0 {
		driver.raw = contents[0]
	} else {
		driver.raw = "{}"
	}

	return nil
}

func (driver *jsonDriver) Set(key string, value any) {
	driver.mutex.Lock()
	defer driver.mutex.Unlock()

	driver.data[key] = value
}

func (driver *jsonDriver) Get(key string) any {
	driver.mutex.RLock()
	defer driver.mutex.RUnlock()

	if v, ok := driver.data[key]; ok {
		return v
	}

	if v := gjson.Get(driver.raw, key); v.Exists() {
		return v.Value()
	}

	return nil
}

func (driver *jsonDriver) Exists(key string) bool {
	driver.mutex.RLock()
	defer driver.mutex.RUnlock()

	if _, ok := driver.data[key]; ok {
		return true
	}

	return gjson.Get(driver.raw, key).Exists()
}

func (driver *jsonDriver) Cast(key string) gocast.Caster {
	return gocast.NewCaster(driver.Get(key))
}
