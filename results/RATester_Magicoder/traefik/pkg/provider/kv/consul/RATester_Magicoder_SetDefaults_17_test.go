package consul

import (
	"fmt"
	"testing"
)

func TestSetDefaults_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	builder := &ProviderBuilder{}
	builder.SetDefaults()

	if builder.Endpoints[0] != "127.0.0.1:8500" {
		t.Errorf("Expected endpoints to be set to '127.0.0.1:8500', but got %s", builder.Endpoints[0])
	}
}
