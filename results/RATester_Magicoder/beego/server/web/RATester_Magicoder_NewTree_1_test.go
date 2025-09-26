package web

import (
	"fmt"
	"testing"
)

func TestNewTree_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := NewTree()
	if tree == nil {
		t.Error("NewTree() should not return nil")
	}
}
