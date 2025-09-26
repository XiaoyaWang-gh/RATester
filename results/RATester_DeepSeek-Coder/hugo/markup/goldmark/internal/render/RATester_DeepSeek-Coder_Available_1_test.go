package render

import (
	"bytes"
	"fmt"
	"testing"
)

func TestAvailable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BufWriter{Buffer: new(bytes.Buffer)}
	expected := maxInt
	actual := b.Available()
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
