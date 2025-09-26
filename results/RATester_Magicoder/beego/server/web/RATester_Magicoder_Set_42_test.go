package web

import (
	"fmt"
	"testing"
)

func TestSet_42(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fd := &FlashData{Data: make(map[string]string)}
	fd.Set("key", "msg")
	if fd.Data["key"] != "msg" {
		t.Errorf("Expected 'msg', got '%s'", fd.Data["key"])
	}
	fd.Set("key", "msg%s", "arg")
	if fd.Data["key"] != "msgarg" {
		t.Errorf("Expected 'msgarg', got '%s'", fd.Data["key"])
	}
}
