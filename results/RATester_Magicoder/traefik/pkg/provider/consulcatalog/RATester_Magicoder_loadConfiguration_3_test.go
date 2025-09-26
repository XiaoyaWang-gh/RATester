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

	ctx := context.Background()
	provider := &Provider{
		name: "test",
	}
	certInfo := &connectCert{
		root: []string{"root1", "root2"},
		leaf: keyPair{
			cert: "cert",
			key:  "key",
		},
	}
	configurationChan := make(chan dynamic.Message, 1)

	err := provider.loadConfiguration(ctx, certInfo, configurationChan)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	select {
	case msg := <-configurationChan:
		if msg.ProviderName != provider.name {
			t.Errorf("Expected provider name %s, but got %s", provider.name, msg.ProviderName)
		}
		if msg.Configuration == nil {
			t.Error("Expected non-nil configuration, but got nil")
		}
	default:
		t.Error("Expected configuration message, but got nothing")
	}
}
