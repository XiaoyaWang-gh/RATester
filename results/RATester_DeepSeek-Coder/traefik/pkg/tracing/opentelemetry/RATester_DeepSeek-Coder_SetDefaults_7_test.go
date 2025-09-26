package opentelemetry

import (
	"fmt"
	"testing"
)

func TestSetDefaults_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := &Config{}
	cfg.SetDefaults()

	if cfg.HTTP == nil {
		t.Errorf("Expected HTTP to be set to a non-nil value")
	}
}
