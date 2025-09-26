package logs

import (
	"fmt"
	"testing"
)

func TestGetLogFuncCallDepth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		loggerFuncCallDepth: 5,
	}

	result := bl.GetLogFuncCallDepth()

	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}
