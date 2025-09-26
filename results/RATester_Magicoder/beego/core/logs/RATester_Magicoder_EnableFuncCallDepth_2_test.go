package logs

import (
	"fmt"
	"testing"
)

func TestEnableFuncCallDepth_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{}
	bl.EnableFuncCallDepth(true)
	if bl.enableFuncCallDepth != true {
		t.Error("Expected EnableFuncCallDepth to set enableFuncCallDepth to true, but it did not.")
	}
}
