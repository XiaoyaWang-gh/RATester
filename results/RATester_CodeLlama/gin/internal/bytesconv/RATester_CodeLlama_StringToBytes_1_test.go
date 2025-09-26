package bytesconv

import (
	"fmt"
	"testing"
)

func TestStringToBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := "hello"
	b := StringToBytes(s)
	if string(b) != s {
		t.Errorf("StringToBytes(%q) = %q, want %q", s, b, s)
	}
}
