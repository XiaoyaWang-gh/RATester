package static

import (
	"fmt"
	"testing"
)

func TestSetDefaults_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tracing := &Tracing{}
	tracing.SetDefaults()

	if tracing.ServiceName != "traefik" {
		t.Errorf("Expected ServiceName to be 'traefik', but got %s", tracing.ServiceName)
	}

	if tracing.SampleRate != 1.0 {
		t.Errorf("Expected SampleRate to be 1.0, but got %f", tracing.SampleRate)
	}

	if tracing.OTLP == nil {
		t.Error("Expected OTLP to be set, but it was nil")
	}
}
