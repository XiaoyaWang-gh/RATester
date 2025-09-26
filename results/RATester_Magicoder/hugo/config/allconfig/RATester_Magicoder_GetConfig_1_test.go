package allconfig

import (
	"fmt"
	"testing"
)

func TestGetConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := &ConfigLanguage{
		config: &Config{
			// Initialize your config here
		},
	}

	result := cfg.GetConfig()

	if result != cfg.config {
		t.Errorf("Expected %v, got %v", cfg.config, result)
	}
}
