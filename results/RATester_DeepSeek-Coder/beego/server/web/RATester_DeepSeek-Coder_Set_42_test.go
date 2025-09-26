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

	fd.Set("key1", "msg1")
	if fd.Data["key1"] != "msg1" {
		t.Errorf("Expected 'msg1', got '%s'", fd.Data["key1"])
	}

	fd.Set("key2", "msg%s", "2")
	if fd.Data["key2"] != "msg2" {
		t.Errorf("Expected 'msg2', got '%s'", fd.Data["key2"])
	}
}
