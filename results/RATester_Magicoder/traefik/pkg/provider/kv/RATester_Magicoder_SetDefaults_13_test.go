package kv

import (
	"fmt"
	"testing"
)

func TestSetDefaults_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{}
	provider.SetDefaults()

	if provider.RootKey != "traefik" {
		t.Errorf("Expected RootKey to be 'traefik', but got %s", provider.RootKey)
	}
}
