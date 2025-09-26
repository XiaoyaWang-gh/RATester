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

	bl := &BeeLogger{
		enableFuncCallDepth: false,
	}

	bl.EnableFuncCallDepth(true)

	if bl.enableFuncCallDepth != true {
		t.Errorf("Expected enableFuncCallDepth to be true, got %v", bl.enableFuncCallDepth)
	}
}
