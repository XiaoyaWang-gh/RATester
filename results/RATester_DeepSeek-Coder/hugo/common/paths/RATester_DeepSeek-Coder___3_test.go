package paths

import (
	"fmt"
	"testing"
)

func Test__3(t *testing.T) {
	t.Run("PathTypeFile", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if PathTypeFile != 0 {
			t.Errorf("Expected PathTypeFile to be 0, got %d", PathTypeFile)
		}
	})

	t.Run("PathTypeContentResource", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if PathTypeContentResource != 1 {
			t.Errorf("Expected PathTypeContentResource to be 1, got %d", PathTypeContentResource)
		}
	})

	t.Run("PathTypeContentSingle", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if PathTypeContentSingle != 2 {
			t.Errorf("Expected PathTypeContentSingle to be 2, got %d", PathTypeContentSingle)
		}
	})

	t.Run("PathTypeLeaf", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if PathTypeLeaf != 3 {
			t.Errorf("Expected PathTypeLeaf to be 3, got %d", PathTypeLeaf)
		}
	})

	t.Run("PathTypeBranch", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if PathTypeBranch != 4 {
			t.Errorf("Expected PathTypeBranch to be 4, got %d", PathTypeBranch)
		}
	})
}
