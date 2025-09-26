package logs

import (
	"fmt"
	"testing"
)

func TestSetLogFuncCallDepth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		loggerFuncCallDepth: 2,
	}

	testDepth := 3
	bl.SetLogFuncCallDepth(testDepth)

	if bl.loggerFuncCallDepth != testDepth {
		t.Errorf("Expected loggerFuncCallDepth to be %d, but got %d", testDepth, bl.loggerFuncCallDepth)
	}
}
