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
		OnDelete:      "",
		OnUpdate:      "",
	}

	expected := "ON DELETECASCADE"
	result := foreign.SetOnDelete("CASCADE")

	if result.OnDelete != expected {
		t.Errorf("Expected OnDelete to be '%s', but got '%s'", expected, result.OnDelete)
	}
}
