package releaser

import (
	"fmt"
	"testing"
)

func TestGitPush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &ReleaseHandler{
		branchVersion: "0.90.0",
		git:           func(args ...string) (string, error) { return "", nil },
	}
	r.gitPush()
}
