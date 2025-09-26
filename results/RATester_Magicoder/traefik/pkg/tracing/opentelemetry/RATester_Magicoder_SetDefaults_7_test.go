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

	config := &Config{}
	config.SetDefaults()

	if config.HTTP == nil {
		t.Error("HTTP should not be nil")
	}
}
