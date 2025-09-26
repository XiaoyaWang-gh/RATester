package nomad

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/nomad/api"
)

func TestBuildConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	provider := &Provider{
		name:      "test",
		namespace: "test",
		client:    &api.Client{},
	}

	items := []item{
		{
			ID:         "1",
			Name:       "service1",
			Namespace:  "test",
			Node:       "node1",
			Datacenter: "dc1",
			Address:    "127.0.0.1",
			Port:       80,
			Tags:       []string{"tag1", "tag2"},
		},
		{
			ID:         "2",
			Name:       "service2",
			Namespace:  "test",
			Node:       "node2",
			Datacenter: "dc2",
			Address:    "127.0.0.2",
			Port:       8080,
			Tags:       []string{"tag3", "tag4"},
		},
	}

	config := provider.buildConfig(ctx, items)

	if config == nil {
		t.Error("Expected non-nil configuration, got nil")
	}

	// Add more assertions as needed
}
