package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestApplyConfigurations_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := &ConfigurationWatcher{
		newConfigs: make(chan dynamic.Configurations),
	}

	go c.applyConfigurations(ctx)

	c.newConfigs <- dynamic.Configurations{}

	cancel()

	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, 1, len(c.newConfigs))
}
