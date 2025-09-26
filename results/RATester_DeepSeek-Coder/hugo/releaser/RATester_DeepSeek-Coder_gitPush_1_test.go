package releaser

import (
	"fmt"
	"testing"
)

func TestGitPush_1(t *testing.T) {
	handler := &ReleaseHandler{
		branchVersion: "v1.0.0",
		skipPush:      true,
		git: func(args ...string) (string, error) {
			if args[0] == "push" && args[1] == "origin" && args[2] == "HEAD" {
				return "", nil
			}
			return "", fmt.Errorf("unexpected git command: %v", args)
		},
	}

	t.Run("when skipPush is true", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		handler.gitPush()
	})

	t.Run("when skipPush is false", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		handler.skipPush = false
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		handler.gitPush()
	})
}
