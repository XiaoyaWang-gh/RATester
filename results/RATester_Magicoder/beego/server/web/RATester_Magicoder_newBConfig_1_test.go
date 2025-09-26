package web

import (
	"fmt"
	"testing"
)

func TestnewBConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := newBConfig()

	if config.AppName != "beego" {
		t.Errorf("Expected AppName to be 'beego', but got '%s'", config.AppName)
	}

	if config.RunMode != "prod" {
		t.Errorf("Expected RunMode to be 'prod', but got '%s'", config.RunMode)
	}

	if config.RouterCaseSensitive != true {
		t.Errorf("Expected RouterCaseSensitive to be true, but got '%t'", config.RouterCaseSensitive)
	}

	if config.RecoverPanic != true {
		t.Errorf("Expected RecoverPanic to be true, but got '%t'", config.RecoverPanic)
	}

	// Add more assertions for other fields as needed
}
