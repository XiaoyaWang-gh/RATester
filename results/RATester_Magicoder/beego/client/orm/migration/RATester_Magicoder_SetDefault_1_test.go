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

	expected := "DEFAULT test_default"
	column.SetDefault("test_default")

	if column.Default != expected {
		t.Errorf("Expected %s, got %s", expected, column.Default)
	}
}
