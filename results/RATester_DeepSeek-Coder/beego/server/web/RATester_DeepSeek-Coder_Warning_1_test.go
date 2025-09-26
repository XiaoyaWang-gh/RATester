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

	fd := &FlashData{
		Data: make(map[string]string),
	}

	msg := "This is a warning message"
	fd.Warning(msg)
	if fd.Data["warning"] != msg {
		t.Errorf("Expected warning message to be '%s', got '%s'", msg, fd.Data["warning"])
	}

	argMsg := "This is a warning message with args: %s"
	arg := "arg"
	fd.Warning(argMsg, arg)
	expectedMsg := fmt.Sprintf(argMsg, arg)
	if fd.Data["warning"] != expectedMsg {
		t.Errorf("Expected warning message to be '%s', got '%s'", expectedMsg, fd.Data["warning"])
	}
}
