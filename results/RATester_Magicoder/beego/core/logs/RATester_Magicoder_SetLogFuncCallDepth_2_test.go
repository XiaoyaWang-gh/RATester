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

	originalDepth := beeLogger.loggerFuncCallDepth
	SetLogFuncCallDepth(10)
	if beeLogger.loggerFuncCallDepth != 10 {
		t.Errorf("Expected loggerFuncCallDepth to be 10, but got %d", beeLogger.loggerFuncCallDepth)
	}
	SetLogFuncCallDepth(originalDepth)
}
