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

func (j *jsonDriver) Load() error {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	if j.data == nil {
		j.data = make(map[string]any)
	}

	// Read json files
	contents := make([]string, 0)
	for _, file := range j.files {
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

		if len(j.files) > 1 {
			contents = append(contents, `"`+fileName+`":`+content)
		} else {
			contents = append(contents, content)
		}
	}

	// Generate big config file
	if len(j.files) > 1 {
		j.raw = "{" + strings.Join(contents, ",") + "}"
	} else if len(j.files) > 0 {
		j.raw = contents[0]
	} else {
		j.raw = "{}"
	}

	return nil
}

func (j *jsonDriver) Set(key string, value any) {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	j.data[key] = value
}

func (j *jsonDriver) Get(key string) any {
	j.mutex.RLock()
	defer j.mutex.RUnlock()

	if v, ok := j.data[key]; ok {
		return v
	}

	if v := gjson.Get(j.raw, key); v.Exists() {
		return v.Value()
	}

	return nil
}

func (j *jsonDriver) Exists(key string) bool {
	j.mutex.RLock()
	defer j.mutex.RUnlock()

	if _, ok := j.data[key]; ok {
		return true
	}

	return gjson.Get(j.raw, key).Exists()
}

func (j *jsonDriver) Cast(key string) gocast.Caster {
	return gocast.NewCaster(j.Get(key))
}
