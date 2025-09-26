package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestOpen_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := filesOnlyFs{
		fs: http.Dir("."),
	}
	f, err := fs.Open("testdata/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "Hello, world!\n" {
		t.Errorf("got %q, want %q", b, "Hello, world!\n")
	}
}
