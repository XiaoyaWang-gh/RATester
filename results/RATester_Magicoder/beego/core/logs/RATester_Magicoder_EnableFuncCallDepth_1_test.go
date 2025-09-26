package logs

import (
	"fmt"
	"testing"
)

func TestEnableFuncCallDepth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	EnableFuncCallDepth(true)
	if !beeLogger.enableFuncCallDepth {
		t.Error("Expected enableFuncCallDepth to be true, but it was false")
	}

	EnableFuncCallDepth(false)
	if beeLogger.enableFuncCallDepth {
		t.Error("Expected enableFuncCallDepth to be false, but it was true")
	}
}
