package releaser

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Testlogln_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rescueStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	logln("Hello, world!")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stderr = rescueStderr

	if string(out) != "Hello, world!\n" {
		t.Errorf("Expected 'Hello, world!', got '%s'", string(out))
	}
}
