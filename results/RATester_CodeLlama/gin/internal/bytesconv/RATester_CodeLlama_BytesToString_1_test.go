package bytesconv

import (
	"fmt"
	"testing"
)

func TestBytesToString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := []byte("hello")
	s := BytesToString(b)
	if s != "hello" {
		t.Errorf("BytesToString(%v) = %v; want %v", b, s, "hello")
	}
}
