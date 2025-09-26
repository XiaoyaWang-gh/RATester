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

	t.Parallel()

	serverTLSConf, clientTLSConf, err := GetTLSConfigs()
	if err != nil {
		t.Fatalf("GetTLSConfigs() failed: %v", err)
	}

	if serverTLSConf == nil || clientTLSConf == nil {
		t.Fatalf("GetTLSConfigs() returned nil for either serverTLSConf or clientTLSConf")
	}

	// Add more specific tests here, for example checking if the certificates are valid, if the private keys are correct, etc.
}
