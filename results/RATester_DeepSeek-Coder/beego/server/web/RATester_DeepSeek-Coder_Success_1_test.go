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

	fd := &FlashData{Data: make(map[string]string)}

	fd.Success("test success message")
	if fd.Data["success"] != "test success message" {
		t.Errorf("Expected 'test success message', got '%s'", fd.Data["success"])
	}

	fd.Success("test %s", "success message")
	if fd.Data["success"] != "test success message" {
		t.Errorf("Expected 'test success message', got '%s'", fd.Data["success"])
	}
}
