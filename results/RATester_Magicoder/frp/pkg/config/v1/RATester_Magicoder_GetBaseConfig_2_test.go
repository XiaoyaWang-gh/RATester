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

	if result.Name != config.Name {
		t.Errorf("Expected Name to be %s, but got %s", config.Name, result.Name)
	}

	if result.Type != config.Type {
		t.Errorf("Expected Type to be %s, but got %s", config.Type, result.Type)
	}
}
