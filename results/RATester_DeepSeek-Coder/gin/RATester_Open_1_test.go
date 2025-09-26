package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestOpen_1(t *testing.T) {
	fs := OnlyFilesFS{http.Dir(".")}

	t.Run("file exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := fs.Open("testdata/testfile.txt")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("file does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := fs.Open("testdata/nonexistentfile.txt")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
