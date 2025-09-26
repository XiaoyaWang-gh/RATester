package releaser

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestLogln_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rescueStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	logln("test")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stderr = rescueStderr

	if !strings.Contains(string(out), "test") {
		t.Errorf("Expected to find 'test', got '%s'", string(out))
	}
}
