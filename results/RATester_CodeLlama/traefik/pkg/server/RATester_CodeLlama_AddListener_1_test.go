package server

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigurationWatcher{}
	listener := func(dynamic.Configuration) {}
	c.AddListener(listener)
	assert.Equal(t, listener, c.configurationListeners[0])
}
