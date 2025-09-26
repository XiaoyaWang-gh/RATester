package echo

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestOpen_1(t *testing.T) {
	t.Run("os.Open", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fs := defaultFS{}
		f, err := fs.Open("testdata/test.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		if string(b) != "Hello, World!\n" {
			t.Errorf("got %q, want %q", b, "Hello, World!\n")
		}
	})
	t.Run("fs.Open", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		fs := defaultFS{fs: os.DirFS("testdata")}
		f, err := fs.Open("test.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		if string(b) != "Hello, World!\n" {
			t.Errorf("got %q, want %q", b, "Hello, World!\n")
		}
	})
}
