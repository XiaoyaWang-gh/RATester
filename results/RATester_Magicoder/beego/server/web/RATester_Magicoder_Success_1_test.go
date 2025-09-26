package web

import (
	"fmt"
	"testing"
)

func TestSuccess_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fd := &FlashData{
		Data: make(map[string]string),
	}

	msg := "test message"
	fd.Success(msg)

	if fd.Data["success"] != msg {
		t.Errorf("Expected success message to be '%s', but got '%s'", msg, fd.Data["success"])
	}

	args := []interface{}{"arg1", "arg2"}
	formattedMsg := fmt.Sprintf(msg, args...)
	fd.Success(msg, args...)

	if fd.Data["success"] != formattedMsg {
		t.Errorf("Expected success message to be '%s', but got '%s'", formattedMsg, fd.Data["success"])
	}
}
