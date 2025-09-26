package testenv

import (
	"fmt"
	"os"
	"testing"
)

func TestMustHaveGoBuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if os.Getenv("GO_GCFLAGS") != "" {
		t.Helper()
		t.Skipf("skipping test: 'go build' not compatible with setting $GO_GCFLAGS")
	}
	if !HasGoBuild() {
		t.Helper()
		t.Skipf("skipping test: 'go build' unavailable: %v", goBuildErr)
	}
}
