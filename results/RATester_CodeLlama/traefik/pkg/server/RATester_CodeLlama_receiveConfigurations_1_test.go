package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestReceiveConfigurations_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := &ConfigurationWatcher{
		allProvidersConfigs: make(chan dynamic.Message),
		newConfigs:          make(chan dynamic.Configurations),
	}

	go c.receiveConfigurations(ctx)

	// TODO: add test cases.
}
