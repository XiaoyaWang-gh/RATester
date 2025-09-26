package session

import (
	"fmt"
	"testing"
)

func TestEncode_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	value := []byte("hello")
	encoded := encode(value)
	if string(encoded) != "aGVsbG8" {
		t.Errorf("encode() = %v, want %v", string(encoded), "aGVsbG8")
	}
}
