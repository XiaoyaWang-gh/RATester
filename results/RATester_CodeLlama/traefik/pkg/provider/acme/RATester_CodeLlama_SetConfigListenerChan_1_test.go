package acme

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestSetConfigListenerChan_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	configFromListenerChan := make(chan dynamic.Configuration)
	p.SetConfigListenerChan(configFromListenerChan)
	assert.Equal(t, configFromListenerChan, p.configFromListenerChan)
}
