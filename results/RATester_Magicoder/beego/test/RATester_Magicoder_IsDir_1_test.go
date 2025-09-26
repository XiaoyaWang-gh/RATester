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

	if fi.IsDir() != false {
		t.Errorf("Expected IsDir() to return false, but got %v", fi.IsDir())
	}
}
