package test

import (
	"fmt"
	"io/fs"
	"testing"
	"time"
)

func TestMode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := bindataFileInfo{
		name:    "test.txt",
		size:    1024,
		mode:    0644,
		modTime: time.Now(),
	}

	expectedMode := fs.FileMode(0644)
	actualMode := fi.Mode()

	if actualMode != expectedMode {
		t.Errorf("Expected Mode: %v, but got: %v", expectedMode, actualMode)
	}
}
