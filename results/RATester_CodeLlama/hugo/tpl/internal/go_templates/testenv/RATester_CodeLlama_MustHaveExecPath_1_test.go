package testenv

import (
	"fmt"
	"testing"
)

func TestMustHaveExecPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	path := "ls"
	MustHaveExecPath(t, path)
}
