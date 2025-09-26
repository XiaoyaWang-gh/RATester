package legacy

import (
	"fmt"
	"testing"
)

func TestGetDefaultServerConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conf := GetDefaultServerConf()

	// Test default values
	if conf.BindAddr != "0.0.0.0" {
		t.Errorf("Expected BindAddr to be '0.0.0.0', but got '%s'", conf.BindAddr)
	}

	// Add more tests as needed
}
