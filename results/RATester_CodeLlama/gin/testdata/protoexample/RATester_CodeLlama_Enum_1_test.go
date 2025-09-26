package protoexample

import (
	"fmt"
	"testing"
)

func TestEnum_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	x := FOO(1)
	p := x.Enum()
	if p == nil {
		t.Errorf("Enum() failed")
	}
	if *p != x {
		t.Errorf("Enum() failed")
	}
}
