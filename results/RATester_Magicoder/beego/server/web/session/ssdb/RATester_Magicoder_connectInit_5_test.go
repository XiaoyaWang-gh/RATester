package ssdb

import (
	"fmt"
	"testing"
)

func TestconnectInit_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{
		Host: "",
		Port: 0,
	}
	err := provider.connectInit()
	if err == nil {
		t.Error("Expected error, got nil")
	}

	provider.Host = "localhost"
	provider.Port = 8888
	err = provider.connectInit()
	if err != nil {
		t.Errorf("Expected nil, got error: %v", err)
	}
}
