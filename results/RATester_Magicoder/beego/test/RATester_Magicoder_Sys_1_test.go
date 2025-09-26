package test

import (
	"fmt"
	"io/fs"
	"testing"
	"time"
)

func TestSys_1(t *testing.T) {
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

	sys := fi.Sys()
	if sys != nil {
		t.Errorf("Expected Sys() to return nil, but got %v", sys)
	}
}
