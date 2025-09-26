package tlstest

import (
	"fmt"
	"testing"
)

func TestGetTLSConfigs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	serverTLSConf, clientTLSConf, err := GetTLSConfigs()
	if err != nil {
		t.Fatalf("Error getting TLS configs: %v", err)
	}

	if serverTLSConf == nil || clientTLSConf == nil {
		t.Fatalf("Expected TLS configs, got nil")
	}

	// Add more specific tests here
}
