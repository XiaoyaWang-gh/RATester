package web

import (
	"fmt"
	"testing"
)

func TestOpen_1(t *testing.T) {
	t.Run("should return a File and nil error when the file exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fs := FileSystem{}
		file, err := fs.Open("testdata/test.txt")
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
		if file == nil {
			t.Error("Expected a File, got nil")
		}
	})

	t.Run("should return an error when the file does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fs := FileSystem{}
		_, err := fs.Open("testdata/non_existing_file.txt")
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	})
}
