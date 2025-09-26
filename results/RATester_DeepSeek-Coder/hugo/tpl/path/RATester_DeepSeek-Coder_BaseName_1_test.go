package path

import (
	"fmt"
	"testing"
)

func TestBaseName_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		path := "/path/to/file.txt"
		expected := "file"
		result, err := ns.BaseName(path)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		path := 1234
		_, err := ns.BaseName(path)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
