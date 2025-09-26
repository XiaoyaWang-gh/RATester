package releaser

import (
	"errors"
	"fmt"
	"testing"
)

func TestgitPush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := &ReleaseHandler{
		branchVersion: "0.90.0",
		step:          1,
		skipPush:      false,
		try:           true,
		git: func(args ...string) (string, error) {
			if args[0] == "push" && args[1] == "origin" && args[2] == "HEAD" {
				return "", nil
			}
			return "", errors.New("push failed")
		},
	}

	handler.gitPush()
}
