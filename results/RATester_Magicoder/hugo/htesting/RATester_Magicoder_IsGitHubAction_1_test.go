package htesting

import (
	"fmt"
	"os"
	"testing"
)

func TestIsGitHubAction_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	os.Setenv("GITHUB_ACTION", "true")
	if !IsGitHubAction() {
		t.Error("Expected true, got false")
	}
	os.Unsetenv("GITHUB_ACTION")
	if IsGitHubAction() {
		t.Error("Expected false, got true")
	}
}
