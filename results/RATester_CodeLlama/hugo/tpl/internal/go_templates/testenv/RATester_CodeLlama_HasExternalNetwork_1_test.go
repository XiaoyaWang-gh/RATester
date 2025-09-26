package testenv

import (
	"fmt"
	"runtime"
	"testing"
)

func TestHasExternalNetwork_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	if runtime.GOOS == "js" || runtime.GOOS == "wasip1" {
		t.Skip("skipping test on js or wasip1")
	}
	if !HasExternalNetwork() {
		t.Error("expected HasExternalNetwork to return true")
	}
}
