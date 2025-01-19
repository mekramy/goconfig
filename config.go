package goconfig

import "github.com/mekramy/gocast"

// Config is an interface that defines methods for loading, setting, getting,
// checking existence, and casting configuration values.
type Config interface {
	// Load loads the configuration from a source.
	Load() error

	// Set sets/override the value for a given key in the configuration.
	Set(key string, value any)

	// Get retrieves the value associated with the given key.
	// on json driver if multiple file passed you must append
	// file name to access setting .e.g. file1.some.field
	Get(key string) any

	// Exists checks if a given key exists in the configuration.
	Exists(key string) bool

	// Cast retrieves the value associated with the given key and returns it
	// as a gocast.Caster, which allows for type-safe casting.
	Cast(key string) gocast.Caster
}

// NewEnv creates a new Config instance that loads configuration from environment variables.
// It accepts an optional list of file paths to load additional environment variables from.
// Returns a Config instance or an error if loading fails.
func NewEnv(files ...string) (Config, error) {
	driver := new(envDriver)
	driver.files = append(driver.files, files...)
	err := driver.Load()
	if err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

// NewJSON creates a new Config instance that loads configuration from JSON files.
// It accepts a list of file paths to load the JSON configuration from.
// Returns a Config instance or an error if loading fails.
func NewJSON(files ...string) (Config, error) {
	driver := new(jsonDriver)
	driver.files = append(driver.files, files...)
	err := driver.Load()
	if err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

// NewMemory creates a new Config instance that loads configuration from an in-memory map.
// It accepts a map of configuration key-value pairs.
// Returns a Config instance or an error if loading fails.
func NewMemory(config map[string]any) (Config, error) {
	driver := new(memoryDriver)
	err := driver.Load()
	if err != nil {
		return nil, err
	}

	for k, v := range config {
		driver.Set(k, v)
	}
	return driver, nil
}
