package etcd

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/provider/kv"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestSetDefaults_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{
		Provider: kv.Provider{
			RootKey:   "traefik",
			Endpoints: []string{"127.0.0.1:2379"},
		},
		TLS: &types.ClientTLS{
			InsecureSkipVerify: true,
		},
		Username: "testuser",
		Password: "testpassword",
	}

	provider.SetDefaults()

	if provider.RootKey != "traefik" {
		t.Errorf("Expected RootKey to be 'traefik', but got %s", provider.RootKey)
	}

	if len(provider.Endpoints) != 1 || provider.Endpoints[0] != "127.0.0.1:2379" {
		t.Errorf("Expected Endpoints to be ['127.0.0.1:2379'], but got %v", provider.Endpoints)
	}

	if provider.TLS.InsecureSkipVerify != true {
		t.Errorf("Expected TLS.InsecureSkipVerify to be true, but got %v", provider.TLS.InsecureSkipVerify)
	}

	if provider.Username != "testuser" {
		t.Errorf("Expected Username to be 'testuser', but got %s", provider.Username)
	}

	if provider.Password != "testpassword" {
		t.Errorf("Expected Password to be 'testpassword', but got %s", provider.Password)
	}
}
