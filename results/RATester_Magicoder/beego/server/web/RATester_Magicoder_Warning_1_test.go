package web

import (
	"fmt"
	"testing"
)

func TestWarning_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fd := &FlashData{Data: make(map[string]string)}
	msg := "test warning"
	fd.Warning(msg)
	if fd.Data["warning"] != msg {
		t.Errorf("Expected warning message to be '%s', but got '%s'", msg, fd.Data["warning"])
	}

	msg = "test warning with args: %s"
	args := "args"
	fd.Warning(msg, args)
	if fd.Data["warning"] != fmt.Sprintf(msg, args) {
		t.Errorf("Expected warning message to be '%s', but got '%s'", fmt.Sprintf(msg, args), fd.Data["warning"])
	}
}
