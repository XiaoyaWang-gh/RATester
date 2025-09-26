package process

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"testing"
)

func TestSetBeforeStopHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &Process{
		cmd:         exec.CommandContext(ctx, "echo", "test"),
		cancel:      cancel,
		errorOutput: &bytes.Buffer{},
		stdOutput:   &bytes.Buffer{},
	}

	testFn := func() {
		// Some test function logic here
	}

	p.SetBeforeStopHandler(testFn)

	if p.beforeStopHandler == nil {
		t.Errorf("Expected beforeStopHandler to be set, but it was nil")
	}
}
