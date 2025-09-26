package utils

import (
	"fmt"
	"testing"
)

func TestAttachFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Email{}
	args := []string{"test.txt"}
	a, err := e.AttachFile(args...)
	if err != nil {
		t.Errorf("AttachFile() error = %v", err)
		return
	}
	if a == nil {
		t.Errorf("AttachFile() a = %v", a)
		return
	}
	if a.Filename != "test.txt" {
		t.Errorf("AttachFile() a.Filename = %v, want %v", a.Filename, "test.txt")
	}
}
