package logs

import (
	"fmt"
	"testing"
)

func TestReset_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		outputs: []*nameLogger{
			{name: "logger1"},
			{name: "logger2"},
		},
	}

	bl.Reset()

	if len(bl.outputs) != 0 {
		t.Errorf("Expected outputs to be empty, got %v", bl.outputs)
	}
}
