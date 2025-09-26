package logs

import (
	"fmt"
	"testing"
)

func TestSetLogFuncCallDepth_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testDepth := 5
	SetLogFuncCallDepth(testDepth)
	if beeLogger.loggerFuncCallDepth != testDepth {
		t.Errorf("Expected SetLogFuncCallDepth to set loggerFuncCallDepth to %d, but got %d", testDepth, beeLogger.loggerFuncCallDepth)
	}
}
