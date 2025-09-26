package orm

import (
	"fmt"
	"os"
	"testing"
)

func TestRunCommand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	os.Args = []string{"orm", "help"}
	RunCommand()

	os.Args = []string{"orm", "unknown"}
	RunCommand()

	os.Args = []string{"orm", "test"}
	RunCommand()
}
