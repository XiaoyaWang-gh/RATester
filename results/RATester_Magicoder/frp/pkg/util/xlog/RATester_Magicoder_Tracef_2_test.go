package xlog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestTracef_2(t *testing.T) {
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

	logger.Tracef("test")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if !strings.Contains(buf.String(), "test") {
		t.Errorf("Tracef did not contain the prefix string")
	}
}
