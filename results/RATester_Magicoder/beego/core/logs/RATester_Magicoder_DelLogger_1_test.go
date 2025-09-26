package logs

import (
	"fmt"
	"testing"
)

func TestDelLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		outputs: []*nameLogger{
			{name: "test1"},
			{name: "test2"},
		},
	}

	err := bl.DelLogger("test1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(bl.outputs) != 1 {
		t.Errorf("Expected 1 output, got %d", len(bl.outputs))
	}
	if bl.outputs[0].name != "test2" {
		t.Errorf("Expected 'test2', got %s", bl.outputs[0].name)
	}

	err = bl.DelLogger("unknown")
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if len(bl.outputs) != 1 {
		t.Errorf("Expected 1 output, got %d", len(bl.outputs))
	}
	if bl.outputs[0].name != "test2" {
		t.Errorf("Expected 'test2', got %s", bl.outputs[0].name)
	}
}
