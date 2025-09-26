package htesting

import (
	"fmt"
	"os"
	"testing"
)

func TestIsGitHubAction_1(t *testing.T) {
	t.Run("TestIsGitHubAction", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("GITHUB_ACTION", "true")
		if !IsGitHubAction() {
			t.Errorf("Expected true, got false")
		}
		os.Unsetenv("GITHUB_ACTION")
		if IsGitHubAction() {
			t.Errorf("Expected false, got true")
		}
	})
}
