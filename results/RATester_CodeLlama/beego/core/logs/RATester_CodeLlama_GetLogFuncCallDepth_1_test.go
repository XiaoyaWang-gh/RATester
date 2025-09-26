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

	bl := &BeeLogger{}
	bl.loggerFuncCallDepth = 1
	if bl.GetLogFuncCallDepth() != 1 {
		t.Errorf("GetLogFuncCallDepth() = %v, want %v", bl.GetLogFuncCallDepth(), 1)
	}
}
