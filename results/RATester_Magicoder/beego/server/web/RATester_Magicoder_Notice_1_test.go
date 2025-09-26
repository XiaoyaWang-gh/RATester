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
	fd.Notice("Hello, %s!", "world")
	if fd.Data["notice"] != "Hello, world!" {
		t.Errorf("Expected 'Hello, world!', got '%s'", fd.Data["notice"])
	}
}
