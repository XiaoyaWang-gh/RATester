package testenv

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMustHaveSymlink_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if runtime.GOOS == "windows" {
		t.Skip("skipping test: cannot make symlinks on windows")
	}
	ok, reason := hasSymlink()
	if !ok {
		t.Skipf("skipping test: cannot make symlinks on %s/%s: %s", runtime.GOOS, runtime.GOARCH, reason)
	}
}
