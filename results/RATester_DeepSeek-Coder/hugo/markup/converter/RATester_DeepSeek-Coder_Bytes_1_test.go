package converter

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := Bytes("test")
	expected := []byte("test")
	result := b.Bytes()
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
