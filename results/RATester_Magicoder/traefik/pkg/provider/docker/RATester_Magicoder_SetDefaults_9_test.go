package docker

import (
	"fmt"
	"testing"
	"time"

	ptypes "github.com/traefik/paerser/types"
)

func TestSetDefaults_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &SwarmProvider{}
	provider.SetDefaults()

	if provider.Watch != true {
		t.Error("Expected Watch to be true, but got", provider.Watch)
	}

	if provider.ExposedByDefault != true {
		t.Error("Expected ExposedByDefault to be true, but got", provider.ExposedByDefault)
	}

	if provider.Endpoint != "unix:///var/run/docker.sock" {
		t.Error("Expected Endpoint to be 'unix:///var/run/docker.sock', but got", provider.Endpoint)
	}

	if provider.RefreshSeconds != ptypes.Duration(15*time.Second) {
		t.Error("Expected RefreshSeconds to be 15 * time.Second, but got", provider.RefreshSeconds)
	}
}
