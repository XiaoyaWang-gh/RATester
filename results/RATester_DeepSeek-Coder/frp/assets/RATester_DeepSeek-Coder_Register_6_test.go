package assets

import (
	"fmt"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestRegister_6(t *testing.T) {
	testFS := fstest.MapFS{
		"static/file.txt": {Data: []byte("Hello, world")},
	}

	t.Run("should register the file system", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		Register(testFS)

		_, err := fs.ReadFile(content, "file.txt")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("should handle error when file system is not found", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		Register(fstest.MapFS{})

		_, err := fs.ReadFile(content, "file.txt")
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
}
