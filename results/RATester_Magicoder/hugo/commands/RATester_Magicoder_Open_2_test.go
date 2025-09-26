package commands

import (
	"fmt"
	"net/http"
	"testing"
)

func TestOpen_2(t *testing.T) {
	fs := filesOnlyFs{http.Dir(".")}

	t.Run("Open existing file", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		f, err := fs.Open("testdata/test.txt")
		if err != nil {
			t.Fatalf("Open() error = %v", err)
		}
		defer f.Close()

		stat, err := f.Stat()
		if err != nil {
			t.Fatalf("Stat() error = %v", err)
		}
		if stat.Name() != "test.txt" {
			t.Errorf("Expected file name to be 'test.txt', got %s", stat.Name())
		}
	})

	t.Run("Open non-existing file", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := fs.Open("testdata/nonexistent.txt")
		if err == nil {
			t.Error("Expected Open() to return an error, but it didn't")
		}
	})
}
