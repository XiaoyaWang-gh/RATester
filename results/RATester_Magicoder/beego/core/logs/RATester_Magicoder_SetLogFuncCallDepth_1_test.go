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

	bl := &BeeLogger{}
	depth := 5
	bl.SetLogFuncCallDepth(depth)
	if bl.loggerFuncCallDepth != depth {
		t.Errorf("Expected loggerFuncCallDepth to be %d, but got %d", depth, bl.loggerFuncCallDepth)
	}
}
