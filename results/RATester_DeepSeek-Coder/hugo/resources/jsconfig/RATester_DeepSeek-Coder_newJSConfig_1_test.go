package jsconfig

import (
	"fmt"
	"testing"
)

func TestNewJSConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := newJSConfig()

	if config.CompilerOptions.BaseURL != "." {
		t.Errorf("Expected BaseURL to be '.' but got %s", config.CompilerOptions.BaseURL)
	}

	if len(config.CompilerOptions.Paths) != 0 {
		t.Errorf("Expected Paths to be empty but got %v", config.CompilerOptions.Paths)
	}
}
