package gin

import (
	"fmt"
	"testing"
)

func TestallocateContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		maxParams: 10,
	}

	ctx := engine.allocateContext(engine.maxParams)

	if ctx == nil {
		t.Error("Expected a non-nil context, but got nil")
	}

	if ctx.engine != engine {
		t.Errorf("Expected context engine to be %v, but got %v", engine, ctx.engine)
	}

	if len(*ctx.params) != 0 {
		t.Errorf("Expected params to be empty, but got %v", *ctx.params)
	}

	if cap(*ctx.params) != int(engine.maxParams) {
		t.Errorf("Expected params capacity to be %v, but got %v", engine.maxParams, cap(*ctx.params))
	}

	if len(*ctx.skippedNodes) != 0 {
		t.Errorf("Expected skippedNodes to be empty, but got %v", *ctx.skippedNodes)
	}

	if cap(*ctx.skippedNodes) != int(engine.maxSections) {
		t.Errorf("Expected skippedNodes capacity to be %v, but got %v", engine.maxSections, cap(*ctx.skippedNodes))
	}
}
