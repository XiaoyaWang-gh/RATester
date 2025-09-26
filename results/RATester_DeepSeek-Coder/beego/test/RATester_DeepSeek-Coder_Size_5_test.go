package test

import (
	"fmt"
	"io/fs"
	"testing"
	"time"
)

func TestSize_5(t *testing.T) {
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

	expected := int64(1024)
	actual := fi.Size()

	if actual != expected {
		t.Errorf("Expected Size() to return %d, but got %d", expected, actual)
	}
}
