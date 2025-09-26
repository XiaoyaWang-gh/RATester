package web

import (
	"fmt"
	"testing"
)

func TestError_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fd := &FlashData{Data: make(map[string]string)}
	msg := "test error"
	fd.Error(msg)
	if fd.Data["error"] != msg {
		t.Errorf("Expected %s, got %s", msg, fd.Data["error"])
	}
}
