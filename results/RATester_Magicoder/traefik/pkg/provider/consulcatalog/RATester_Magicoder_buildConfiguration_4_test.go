package consulcatalog

import (
	"context"
	"fmt"
	"testing"
	"text/template"

	"github.com/hashicorp/consul/api"
)

func TestBuildConfiguration_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()

	provider := &Provider{
		name:              "test",
		namespace:         "default",
		client:            &api.Client{},
		defaultRuleTpl:    template.New("test"),
		certChan:          make(chan *connectCert),
		watchServicesChan: make(chan struct{}),
	}

	items := []itemData{
		{
			ID:      "1",
			Node:    "node1",
			Name:    "service1",
			Address: "127.0.0.1",
			Port:    "8080",
			Labels:  map[string]string{"traefik.enable": "true"},
		},
		{
			ID:      "2",
			Node:    "node2",
			Name:    "service2",
			Address: "127.0.0.1",
			Port:    "8081",
			Labels:  map[string]string{"traefik.enable": "false"},
		},
	}

	config := provider.buildConfiguration(ctx, items, &connectCert{})

	if config == nil {
		t.Error("Expected configuration, got nil")
	}

	if len(config.HTTP.Routers) != 1 {
		t.Errorf("Expected 1 router, got %d", len(config.HTTP.Routers))
	}

	if _, ok := config.HTTP.Routers["service1-node1-1"]; !ok {
		t.Error("Expected router for service1, got none")
	}

	if _, ok := config.HTTP.Routers["service2-node2-2"]; ok {
		t.Error("Expected no router for service2, got one")
	}
}
