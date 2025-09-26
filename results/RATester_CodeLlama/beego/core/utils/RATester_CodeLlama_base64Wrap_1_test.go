package utils

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBase64Wrap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var w bytes.Buffer
	var b []byte
	base64Wrap(&w, b)
	if got, want := w.String(), ""; got != want {
		t.Errorf("base64Wrap() = %q, want %q", got, want)
	}
}
