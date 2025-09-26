package releaser

import (
	"fmt"
	"testing"
)

func TestGitPull_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := &ReleaseHandler{
		git: func(args ...string) (string, error) {
			if args[0] != "pull" || args[1] != "origin" || args[2] != "HEAD" {
				t.Errorf("Expected 'pull origin HEAD', got '%s %s %s'", args[0], args[1], args[2])
			}
			return "", nil
		},
	}

	handler.gitPull()
}
