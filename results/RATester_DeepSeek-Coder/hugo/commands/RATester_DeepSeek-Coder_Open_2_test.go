package commands

import (
	"fmt"
	"net/http"
	"testing"
)

func TestOpen_2(t *testing.T) {
	fs := filesOnlyFs{http.Dir(".")}

	t.Run("file exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		f, err := fs.Open("testdata/test.txt")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer f.Close()

		stat, err := f.Stat()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !stat.Mode().IsRegular() {
			t.Errorf("expected a regular file")
		}
	})

	t.Run("file does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := fs.Open("testdata/doesnotexist.txt")
		if err == nil {
			t.Errorf("expected an error")
		}
	})
}
