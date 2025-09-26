package migration

import (
	"fmt"
	"testing"
)

func TestSetOnUpdate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	foreign := &Foreign{
		ForeignTable:  "test_table",
		ForeignColumn: "test_column",
	}

	updatedForeign := foreign.SetOnUpdate(" CASCADE")

	if updatedForeign.OnUpdate != "ON UPDATE CASCADE" {
		t.Errorf("Expected OnUpdate to be 'ON UPDATE CASCADE', but got '%s'", updatedForeign.OnUpdate)
	}
}
