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

	column := &Column{}
	column.SetAuto(true)
	if column.Inc != "auto_increment" {
		t.Errorf("Expected 'auto_increment', but got '%s'", column.Inc)
	}
}
