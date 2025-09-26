package hugolib

import (
	"fmt"
	"testing"
)

func TestDefaultConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := DefaultConfig()

	if config == nil {
		t.Error("Expected config to be not nil")
	}

	if config.BaseURL != "" {
		t.Errorf("Expected BaseURL to be empty, got %s", config.BaseURL)
	}
}
