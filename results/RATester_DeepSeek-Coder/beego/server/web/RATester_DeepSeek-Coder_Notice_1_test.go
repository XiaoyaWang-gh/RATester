package web

import (
	"fmt"
	"testing"
)

func TestNotice_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fd := &FlashData{Data: make(map[string]string)}

	fd.Notice("test")
	if fd.Data["notice"] != "test" {
		t.Errorf("Expected 'test', got '%s'", fd.Data["notice"])
	}

	fd.Notice("test %s", "args")
	if fd.Data["notice"] != "test args" {
		t.Errorf("Expected 'test args', got '%s'", fd.Data["notice"])
	}
}
