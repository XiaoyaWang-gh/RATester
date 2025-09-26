package logs

import (
	"fmt"
	"testing"
)

func TestDelLogger_1(t *testing.T) {
	bl := &BeeLogger{
		outputs: []*nameLogger{
			{name: "test1"},
			{name: "test2"},
			{name: "test3"},
		},
	}

	t.Run("ExistingAdapter", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := bl.DelLogger("test2")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(bl.outputs) != 2 {
			t.Errorf("Expected 2 outputs, got %d", len(bl.outputs))
		}
		if bl.outputs[0].name != "test1" || bl.outputs[1].name != "test3" {
			t.Errorf("Expected 'test1' and 'test3', got %s and %s", bl.outputs[0].name, bl.outputs[1].name)
		}
	})

	t.Run("NonExistingAdapter", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := bl.DelLogger("test4")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if len(bl.outputs) != 2 {
			t.Errorf("Expected 2 outputs, got %d", len(bl.outputs))
		}
		if bl.outputs[0].name != "test1" || bl.outputs[1].name != "test3" {
			t.Errorf("Expected 'test1' and 'test3', got %s and %s", bl.outputs[0].name, bl.outputs[1].name)
		}
	})
}
