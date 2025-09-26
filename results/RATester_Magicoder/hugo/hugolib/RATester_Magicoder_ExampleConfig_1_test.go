package hugolib

import (
	"fmt"
	"testing"
)

func TestExampleConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config, err := ExampleConfig()
	if err != nil {
		t.Errorf("Error in ExampleConfig: %v", err)
	}

	// Test the baseURL
	if config.BaseURL != "https://example.com/" {
		t.Errorf("Expected BaseURL to be 'https://example.com/', but got '%s'", config.BaseURL)
	}

	// Add more tests as needed
}
