package testenv

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestGoTool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Test case 1: Go tool exists
	exec.Command("go").Run()
	goToolPath, goToolErr := GoTool()
	if goToolErr != nil {
		t.Errorf("GoTool() returned an error: %v", goToolErr)
	}
	if goToolPath == "" {
		t.Errorf("GoTool() returned an empty string")
	}

	// Test case 2: Go tool does not exist
	exec.Command("rm", "go").Run()
	goToolPath, goToolErr = GoTool()
	if goToolErr == nil {
		t.Errorf("GoTool() did not return an error when go tool does not exist")
	}
	if goToolPath != "" {
		t.Errorf("GoTool() returned a non-empty string when go tool does not exist")
	}
}
