package utils

import (
	"fmt"
	"testing"
)

func TestQpEscape_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dest := make([]byte, 3)
	c := byte(0x1)
	qpEscape(dest, c)
	if dest[0] != '=' || dest[1] != '0' || dest[2] != '1' {
		t.Errorf("qpEscape(dest, c) = %v, want %v", dest, []byte{'=', '0', '1'})
	}
}
