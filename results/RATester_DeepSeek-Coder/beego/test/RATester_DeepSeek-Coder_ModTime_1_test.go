package test

import (
	"fmt"
	"io/fs"
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
		mode:    fs.FileMode(0644),
		modTime: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	expectedModTime := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	actualModTime := fi.ModTime()

	if actualModTime != expectedModTime {
		t.Errorf("Expected ModTime to be %v, but got %v", expectedModTime, actualModTime)
	}
}
