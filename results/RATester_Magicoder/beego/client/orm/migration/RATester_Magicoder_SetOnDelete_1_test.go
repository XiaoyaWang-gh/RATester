package migration

import (
	"fmt"
	"testing"
)

func TestSetOnDelete_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	foreign := &Foreign{
		ForeignTable:  "test_table",
		ForeignColumn: "test_column",
	}

	foreign.SetOnDelete(" CASCADE")

	if foreign.OnDelete != "ON DELETE CASCADE" {
		t.Errorf("Expected OnDelete to be 'ON DELETE CASCADE', but got '%s'", foreign.OnDelete)
	}
}
