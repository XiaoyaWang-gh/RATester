package fiber

import (
	"fmt"
	"runtime"
	"testing"
)

func TestDummyCmd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	if runtime.GOOS == "windows" {
		t.Skip("skip on windows")
	}
	if dummyChildCmd.Load() != nil {
		t.Fatal("dummyChildCmd should be nil")
	}
	dummyChildCmd.Store("go")
	if dummyCmd().Path != "go" {
		t.Fatal("dummyCmd should be go")
	}
	dummyChildCmd.Store("")
	if dummyCmd().Path != "go" {
		t.Fatal("dummyCmd should be go")
	}
}
