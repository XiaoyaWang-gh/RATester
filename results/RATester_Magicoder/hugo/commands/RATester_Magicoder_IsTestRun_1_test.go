package commands

import (
	"fmt"
	"os"
	"testing"
)

func TestIsTestRun_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	os.Setenv("HUGO_TESTRUN", "true")
	defer os.Unsetenv("HUGO_TESTRUN")

	root := &rootCommand{}

	if !root.IsTestRun() {
		t.Error("Expected IsTestRun to return true when HUGO_TESTRUN is set")
	}
}
