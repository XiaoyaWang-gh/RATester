package migration

import (
	"fmt"
	"testing"
)

func TestSetDefault_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	column := &Column{
		Name: "test_column",
	}

	expectedDefault := "DEFAULT test_default"
	result := column.SetDefault("test_default")

	if result.Default != expectedDefault {
		t.Errorf("Expected default value to be '%s', but got '%s'", expectedDefault, result.Default)
	}
}
