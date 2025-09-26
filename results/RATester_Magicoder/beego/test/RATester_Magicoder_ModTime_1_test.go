package test

import (
	"fmt"
	"testing"
	"time"
)

func TestModTime_1(t *testing.T) {
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

	expectedModTime := fi.modTime
	actualModTime := fi.ModTime()

	if !expectedModTime.Equal(actualModTime) {
		t.Errorf("Expected ModTime: %v, but got: %v", expectedModTime, actualModTime)
	}
}
