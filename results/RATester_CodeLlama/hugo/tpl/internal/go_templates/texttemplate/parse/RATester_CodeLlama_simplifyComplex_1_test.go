package parse

import (
	"fmt"
	"testing"
)

func TestSimplifyComplex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &NumberNode{
		Complex128: complex(1, 2),
	}
	n.simplifyComplex()
	if !n.IsFloat {
		t.Errorf("IsFloat should be true")
	}
	if n.Float64 != 1 {
		t.Errorf("Float64 should be 1, got %v", n.Float64)
	}
	if !n.IsInt {
		t.Errorf("IsInt should be true")
	}
	if n.Int64 != 1 {
		t.Errorf("Int64 should be 1, got %v", n.Int64)
	}
	if !n.IsUint {
		t.Errorf("IsUint should be true")
	}
	if n.Uint64 != 1 {
		t.Errorf("Uint64 should be 1, got %v", n.Uint64)
	}
}
