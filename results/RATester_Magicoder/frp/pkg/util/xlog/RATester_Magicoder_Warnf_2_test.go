package xlog

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestWarnf_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := &Logger{}
	logger.prefixString = "test"

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger.Warnf("test")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old

	if !strings.Contains(string(out), "test") {
		t.Errorf("Expected output to contain 'test', but got '%s'", string(out))
	}
}
