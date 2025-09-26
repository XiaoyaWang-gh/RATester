package orm

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestprintHelp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printHelp("Error message")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	expected := "Error message\n" + `orm command usage:

    syncdb     - auto create tables
    sqlall     - print sql of create tables
    help       - print this help
`
	if string(out) != expected {
		t.Errorf("Expected: %s, Got: %s", expected, string(out))
	}
}
