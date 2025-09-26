package echo

import (
	"fmt"
	"testing"
)

func TestOpen_1(t *testing.T) {
	t.Run("should return error when file does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fs := defaultFS{}
		_, err := fs.Open("non-existing-file")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("should return file when file exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fs := defaultFS{}
		file, err := fs.Open("existing-file")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if file == nil {
			t.Errorf("Expected file, got nil")
		}
	})
}
