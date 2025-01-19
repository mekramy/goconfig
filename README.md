# GoConfig

GoConfig is a Go package that provides a flexible and easy-to-use configuration management system. It supports loading configuration from various sources such as environment variables, JSON files, and in-memory maps. The package defines a `Config` interface with methods for loading, setting, getting, checking existence, and casting configuration values.

## Installation

To install GoConfig, use the following command:

```sh
go get github.com/mekramy/goconfig
```

### Interface Methods

The `Config` interface provides the following methods:

- `Load() error`: Loads configuration from the specified source.
- `Set(key string, value interface{})`: Sets a configuration value for the given key. If the configuration key exists in the config files, this method overrides it.
- `Get(key string) interface{}`: Retrieves the configuration value for the given key.
- `Exists(key string) bool`: Checks if a configuration value exists for the given key.
- `Cast(key string) (gocast.Caster, error)`: Casts the configuration value to the specified type.

### Constructor Methods

#### NewEnv

Creates a new Config instance that loads configuration from `.env` files to environment variables and reads config from OS environment variables.

#### NewJSON

Creates a new Config instance that loads configuration from JSON files. If multiple files are passed, the file name must be specified as part of the config key.

#### NewMemory

Creates a new Config instance that loads configuration from an in-memory map.

### Example

```go
package main

import (
    "fmt"
    "github.com/mekramy/goconfig"
)

func main() {
    // Load configuration from an env file
    env := goconfig.NewEnv(".env")
    if err != nil {
        fmt.Println("Error loading config:", err)
        return
    }

    // Override value
    env.Set("DATABASE_PASSWORD", "my-secret-password")

    // Check value
    if env.Exists("REDIS") {
        fmt.Println("REDIS config is defined")
    }

    // Cast value
    if onProd := env.Cast("PROD").BoolSafe(false); onProd {
        fmt.Println("App runs in production mode")
    }

    // Load configuration from multiple JSON files
    // If config is loaded from a single file, you must
    // enter the config path without the file name
    //
    // app.json content
    // {
    //      "meta": {
    //          "name": "test",
    //          "version": "v2.3.0"
    //      }
    // }
    //
    // http.json
    // {
    //      "port": 8080,
    //      "ssl": false,
    //      "domain": "my-app.com"
    // }

    json, err := goconfig.NewJSON("app.json", "http.json")
    if err != nil {
        fmt.Println("Error loading config:", err)
        return
    }

    // Access config with fallback
    port := json.Cast("http.port").IntSafe(8888) // use 8888 if http.port does not exist
    version := json.Cast("app.meta.version").StringSafe("v0.0.0")
}
```
