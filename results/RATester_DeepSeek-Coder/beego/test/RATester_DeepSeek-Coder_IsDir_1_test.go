package test

import (
	"fmt"
	"io/fs"
	"testing"
	"time"
)

func TestIsDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := bindataFileInfo{
		name:    "test.txt",
		size:    1024,
		mode:    fs.FileMode(0644),
		modTime: time.Now(),
	}

	result := fi.IsDir()
	if result != false {
		t.Errorf("Expected IsDir() to return false, got %v", result)
	}
}
