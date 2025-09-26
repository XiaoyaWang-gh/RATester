package consulcatalog

import (
	"fmt"
	"testing"
	"time"

	ptypes "github.com/traefik/paerser/types"
)

func TestSetDefaults_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &Configuration{}
	config.SetDefaults()

	if config.Endpoint == nil {
		t.Error("Endpoint should not be nil")
	}

	if config.RefreshInterval != ptypes.Duration(15*time.Second) {
		t.Error("RefreshInterval should be 15 seconds")
	}

	if config.Prefix != "traefik" {
		t.Error("Prefix should be 'traefik'")
	}

	if !config.ExposedByDefault {
		t.Error("ExposedByDefault should be true")
	}

	if config.DefaultRule != defaultTemplateRule {
		t.Error("DefaultRule should be the default template rule")
	}

	if config.ServiceName != "traefik" {
		t.Error("ServiceName should be 'traefik'")
	}

	if len(config.StrictChecks) != 0 {
		t.Error("StrictChecks should be empty")
	}
}
