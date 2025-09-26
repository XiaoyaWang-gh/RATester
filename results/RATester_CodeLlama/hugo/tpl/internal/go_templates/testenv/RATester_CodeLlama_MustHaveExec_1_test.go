package testenv

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMustHaveExec_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if runtime.GOOS == "windows" {
		t.Skip("skipping test on windows")
	}
	MustHaveExec(t)
}
