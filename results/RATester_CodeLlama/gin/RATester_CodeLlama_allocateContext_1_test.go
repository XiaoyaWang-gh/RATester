package gin

import (
	"fmt"
	"testing"
)

func TestAllocateContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	maxParams := uint16(10)
	engine := &Engine{}
	context := engine.allocateContext(maxParams)
	if context.engine != engine {
		t.Errorf("engine is %v, want %v", context.engine, engine)
	}
	if context.params == nil {
		t.Errorf("params is nil, want not nil")
	}
	if context.skippedNodes == nil {
		t.Errorf("skippedNodes is nil, want not nil")
	}
}
