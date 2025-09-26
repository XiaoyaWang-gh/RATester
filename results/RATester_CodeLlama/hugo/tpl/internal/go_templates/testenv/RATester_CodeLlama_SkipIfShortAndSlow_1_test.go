package testenv

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSkipIfShortAndSlow_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if testing.Short() && runtime.GOARCH == "amd64" {
		t.Helper()
		t.Skipf("skipping test in -short mode on %s", runtime.GOARCH)
	}
}
