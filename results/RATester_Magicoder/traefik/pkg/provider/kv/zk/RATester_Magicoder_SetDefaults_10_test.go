package zk

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/provider/kv"
)

func TestSetDefaults_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{
		Provider: kv.Provider{
			RootKey:   "test",
			Endpoints: []string{"127.0.0.1:2181"},
		},
		Username: "test_user",
		Password: "test_password",
	}

	provider.SetDefaults()

	if provider.Endpoints[0] != "127.0.0.1:2181" {
		t.Errorf("Expected Endpoints to be unchanged, but got %v", provider.Endpoints)
	}

	if provider.Username != "test_user" {
		t.Errorf("Expected Username to be unchanged, but got %v", provider.Username)
	}

	if provider.Password != "test_password" {
		t.Errorf("Expected Password to be unchanged, but got %v", provider.Password)
	}
}
