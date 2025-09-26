package http

import (
	"fmt"
	"testing"
)

func TestPathPrefix_1(t *testing.T) {
	tree := &matchersTree{}

	t.Run("valid path", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := pathPrefix(tree, "/valid/path")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("invalid path", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		err := pathPrefix(tree, "invalid/path")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
