package consulcatalog

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestWatchConnectTLS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	provider := &Provider{
		client: &api.Client{},
	}

	err := provider.watchConnectTLS(ctx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
