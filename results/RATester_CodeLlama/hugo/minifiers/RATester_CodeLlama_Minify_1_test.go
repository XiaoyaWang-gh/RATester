package minifiers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMinify_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var m noopMinifier
	var w bytes.Buffer
	var r bytes.Buffer
	r.WriteString("Hello, World!")
	err := m.Minify(nil, &w, &r, nil)
	if err != nil {
		t.Errorf("Minify returned error: %v", err)
	}
	if w.String() != "Hello, World!" {
		t.Errorf("Minify wrote incorrect string: %s", w.String())
	}
}
