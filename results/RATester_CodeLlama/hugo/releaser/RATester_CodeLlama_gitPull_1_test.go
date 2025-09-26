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

	r := &ReleaseHandler{
		git: func(args ...string) (string, error) {
			return "git pull origin HEAD", nil
		},
	}
	r.gitPull()
}
