package testenv

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMustHaveLink_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if runtime.GOOS == "windows" {
		t.Skip("skipping test: hardlinks are not supported on windows")
	}
	if !HasLink() {
		t.Error("hardlinks are not supported on this platform")
	}
}
