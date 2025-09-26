package migration

import (
	"fmt"
	"testing"
)

func TestSetAuto_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	column := &Column{
		Name: "test_column",
	}

	column.SetAuto(true)

	if column.Inc != "auto_increment" {
		t.Errorf("Expected 'auto_increment', got '%s'", column.Inc)
	}

	column.SetAuto(false)

	if column.Inc != "" {
		t.Errorf("Expected '', got '%s'", column.Inc)
	}
}
