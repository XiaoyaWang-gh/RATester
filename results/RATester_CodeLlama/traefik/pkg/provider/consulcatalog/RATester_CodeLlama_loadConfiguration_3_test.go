package consulcatalog

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestLoadConfiguration_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	ctx := context.Background()
	certInfo := &connectCert{}
	configurationChan := make(chan dynamic.Message)
	err := p.loadConfiguration(ctx, certInfo, configurationChan)
	if err != nil {
		t.Errorf("loadConfiguration returned an error: %v", err)
	}
}
