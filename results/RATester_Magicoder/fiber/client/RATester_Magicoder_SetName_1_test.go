package client

import (
	"fmt"
	"testing"
)

func TestSetName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file := &File{}
	file.SetName("test.txt")

	if file.name != "test.txt" {
		t.Errorf("Expected name to be 'test.txt', but got '%s'", file.name)
	}
}
