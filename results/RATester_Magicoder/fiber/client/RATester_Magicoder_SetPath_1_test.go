package client

import (
	"fmt"
	"testing"
)

func TestSetPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file := &File{}
	path := "test.txt"
	file.SetPath(path)
	if file.path != path {
		t.Errorf("Expected path to be %s, but got %s", path, file.path)
	}
}
