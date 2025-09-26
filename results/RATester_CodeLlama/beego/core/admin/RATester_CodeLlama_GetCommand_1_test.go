package admin

import (
	"fmt"
	"testing"
)

func TestGetCommand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	module := "module1"
	cmdName := "cmd1"
	cmd := GetCommand(module, cmdName)
	if cmd == nil {
		t.Errorf("GetCommand failed, module: %s, cmdName: %s", module, cmdName)
	}
}
