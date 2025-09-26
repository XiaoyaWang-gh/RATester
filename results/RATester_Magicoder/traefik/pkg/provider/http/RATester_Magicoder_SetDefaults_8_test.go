package http

import (
	"fmt"
	"testing"
	"time"

	ptypes "github.com/traefik/paerser/types"
)

func TestSetDefaults_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{}
	provider.SetDefaults()

	if provider.PollInterval != ptypes.Duration(5*time.Second) {
		t.Errorf("Expected PollInterval to be 5 seconds, but got %v", provider.PollInterval)
	}

	if provider.PollTimeout != ptypes.Duration(5*time.Second) {
		t.Errorf("Expected PollTimeout to be 5 seconds, but got %v", provider.PollTimeout)
	}
}
