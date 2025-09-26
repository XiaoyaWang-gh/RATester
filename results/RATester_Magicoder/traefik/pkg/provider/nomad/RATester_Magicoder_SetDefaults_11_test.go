package nomad

import (
	"fmt"
	"testing"
)

func TestSetDefaults_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{
		Configuration: Configuration{
			DefaultRule: "default",
		},
	}

	provider.SetDefaults()

	if provider.Configuration.DefaultRule != "default" {
		t.Errorf("Expected DefaultRule to be 'default', but got %s", provider.Configuration.DefaultRule)
	}
}
