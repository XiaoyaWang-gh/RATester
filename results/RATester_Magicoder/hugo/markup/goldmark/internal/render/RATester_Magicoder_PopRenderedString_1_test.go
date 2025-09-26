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
		BufWriter: &BufWriter{
			Buffer: bytes.NewBuffer([]byte("test string")),
		},
	}

	expected := "test string"
	actual := ctx.PopRenderedString()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
