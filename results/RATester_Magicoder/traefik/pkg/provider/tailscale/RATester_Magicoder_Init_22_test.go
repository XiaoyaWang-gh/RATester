package tailscale

import (
	"fmt"
	"testing"
)

func TestInit_22(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{}
	err := provider.Init()
	if err != nil {
		t.Errorf("Init() failed: %v", err)
	}

	if provider.dynConfigs == nil {
		t.Error("dynConfigs is nil after Init()")
	}

	if provider.certByDomain == nil {
		t.Error("certByDomain is nil after Init()")
	}
}
