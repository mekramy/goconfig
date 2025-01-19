package goconfig_test

import (
	"testing"

	"github.com/mekramy/goconfig"
)

func TestNewEnv(t *testing.T) {
	config, err := goconfig.NewEnv("test.env")
	if err != nil {
		t.Fatal(err)
	}

	v, err := config.Cast("APP_TITLE").String()
	if err != nil {
		t.Fatal(err)
	}

	if v != "My App" {
		t.Errorf(`Failed check APP_TITLE == "My App"`)
	}
}

func TestNewJSON(t *testing.T) {
	config, err := goconfig.NewJSON("test.json")
	if err != nil {
		t.Fatal(err)
	}

	v, err := config.Cast("app.title").String()
	if err != nil {
		t.Fatal(err)
	}

	if v != "My App" {
		t.Errorf(`Failed check app.title == "My App"`)
	}
}

func TestNewMemory(t *testing.T) {
	config, err := goconfig.NewMemory(map[string]any{"title": "My App"})
	if err != nil {
		t.Fatal(err)
	}

	v, err := config.Cast("title").String()
	if err != nil {
		t.Fatal(err)
	}

	if v != "My App" {
		t.Errorf(`Failed check title == "My App"`)
	}
}
