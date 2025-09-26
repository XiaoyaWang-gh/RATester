package testenv

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMustHaveParallelism_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if !HasParallelism() {
		t.Skipf("skipping test: no parallelism available on %s/%s", runtime.GOOS, runtime.GOARCH)
	}
}
