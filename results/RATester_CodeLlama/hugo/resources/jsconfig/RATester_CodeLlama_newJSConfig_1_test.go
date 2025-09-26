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

	// Arrange
	// Act
	config := newJSConfig()
	// Assert
	if config == nil {
		t.Errorf("newJSConfig() returned nil")
	}
	if config.CompilerOptions.BaseURL != "." {
		t.Errorf("newJSConfig() returned CompilerOptions.BaseURL != \".\"")
	}
	if len(config.CompilerOptions.Paths) != 0 {
		t.Errorf("newJSConfig() returned CompilerOptions.Paths != nil")
	}
}
