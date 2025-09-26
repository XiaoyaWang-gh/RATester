package types

import (
	"fmt"
	"testing"
)

func TestSetDefaults_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	otelHTTP := &OtelHTTP{}
	otelHTTP.SetDefaults()

	if otelHTTP.Endpoint != "https://localhost:4318" {
		t.Errorf("Expected endpoint to be 'https://localhost:4318', but got '%s'", otelHTTP.Endpoint)
	}
}
