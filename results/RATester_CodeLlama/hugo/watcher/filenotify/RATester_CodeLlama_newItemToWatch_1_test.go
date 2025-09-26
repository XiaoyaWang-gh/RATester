package filenotify

import (
	"fmt"
	"testing"
)

func TestNewItemToWatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	filename := "testdata/test.txt"
	// Act
	item, err := newItemToWatch(filename)
	// Assert
	if err != nil {
		t.Errorf("newItemToWatch() error = %v", err)
		return
	}
	if item == nil {
		t.Errorf("newItemToWatch() item = nil")
		return
	}
	if item.filename != filename {
		t.Errorf("newItemToWatch() item.filename = %v, want %v", item.filename, filename)
	}
	if item.left == nil {
		t.Errorf("newItemToWatch() item.left = nil")
	}
	if item.right != nil {
		t.Errorf("newItemToWatch() item.right = %v, want nil", item.right)
	}
}
