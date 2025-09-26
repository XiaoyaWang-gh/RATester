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
		OnDelete:      "ON DELETE",
		OnUpdate:      "ON UPDATE",
	}

	update := " CASCADE"
	expected := "ON UPDATE" + update

	foreign.SetOnUpdate(update)

	if foreign.OnUpdate != expected {
		t.Errorf("Expected OnUpdate to be '%s', but got '%s'", expected, foreign.OnUpdate)
	}
}
