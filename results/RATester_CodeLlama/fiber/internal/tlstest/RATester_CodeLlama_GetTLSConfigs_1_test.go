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
		t.Fatal(err)
	}

	if serverTLSConf == nil {
		t.Fatal("serverTLSConf is nil")
	}

	if clientTLSConf == nil {
		t.Fatal("clientTLSConf is nil")
	}
}
