package nomad

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestLoadConfiguration_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	configurationC := make(chan dynamic.Message)
	p := &Provider{}
	err := p.loadConfiguration(ctx, configurationC)
	if err != nil {
		t.Errorf("loadConfiguration returned an error: %v", err)
	}
}
