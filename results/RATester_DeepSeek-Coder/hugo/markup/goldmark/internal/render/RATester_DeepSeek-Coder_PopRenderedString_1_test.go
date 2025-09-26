package render

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPopRenderedString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		BufWriter: &BufWriter{Buffer: &bytes.Buffer{}},
		positions: []int{10},
	}
	ctx.Buffer.WriteString("Hello, World")

	expected := "Hello, World"
	result := ctx.PopRenderedString()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}

	if ctx.Buffer.String() != "" {
		t.Errorf("Expected buffer to be empty, but got '%s'", ctx.Buffer.String())
	}
}
