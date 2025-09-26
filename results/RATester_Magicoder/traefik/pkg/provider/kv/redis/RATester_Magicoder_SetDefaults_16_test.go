package redis

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/provider/kv"
)

func TestSetDefaults_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := Provider{
		Provider: kv.Provider{
			RootKey:   "test",
			Endpoints: []string{"127.0.0.1:6379"},
		},
	}

	provider.SetDefaults()

	if provider.Endpoints[0] != "127.0.0.1:6379" {
		t.Errorf("Expected Endpoints to be '127.0.0.1:6379', but got '%s'", provider.Endpoints[0])
	}
}
