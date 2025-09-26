package types

import (
	"fmt"
	"testing"
)

func TestNewBool_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := true
	var p *bool
	p = NewBool(b)
	if *p != b {
		t.Errorf("NewBool(%v) = %v, want %v", b, *p, b)
	}
}
