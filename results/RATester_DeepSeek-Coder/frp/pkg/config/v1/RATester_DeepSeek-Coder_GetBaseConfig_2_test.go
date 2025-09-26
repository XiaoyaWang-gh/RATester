package v1

import (
	"fmt"
	"testing"
)

func TestGetBaseConfig_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &VisitorBaseConfig{
		Name: "test",
		Type: "test",
	}

	result := config.GetBaseConfig()

	if result.Name != "test" {
		t.Errorf("Expected Name to be 'test', got %s", result.Name)
	}

	if result.Type != "test" {
		t.Errorf("Expected Type to be 'test', got %s", result.Type)
	}
}
