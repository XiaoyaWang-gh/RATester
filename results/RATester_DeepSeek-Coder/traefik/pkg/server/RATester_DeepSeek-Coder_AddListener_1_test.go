package server

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	watcher := &ConfigurationWatcher{
		configurationListeners: make([]func(dynamic.Configuration), 0),
	}

	listener := func(config dynamic.Configuration) {
		// do something with config
	}

	watcher.AddListener(listener)

	if len(watcher.configurationListeners) != 1 {
		t.Errorf("Expected 1 listener, got %d", len(watcher.configurationListeners))
	}
}
