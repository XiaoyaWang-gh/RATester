package testenv

import (
	"fmt"
	"runtime"
	"testing"
)

func TestParallelOn64Bit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	is64Bit := runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
	if !is64Bit {
		t.Skip("Test is only for 64-bit architectures")
	}

	t.Parallel()

	// Test code here
}
