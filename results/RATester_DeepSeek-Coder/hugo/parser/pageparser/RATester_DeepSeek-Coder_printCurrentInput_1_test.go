package pageparser

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestPrintCurrentInput_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{
		input: []byte("test input"),
		pos:   5,
	}

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	l.printCurrentInput()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	expected := "input[5:]: \"input\""
	if string(out) != expected {
		t.Errorf("Expected %q, got %q", expected, string(out))
	}
}
