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
	result := x.Enum()
	if *result != x {
		t.Errorf("Expected %v, got %v", x, *result)
	}
}
