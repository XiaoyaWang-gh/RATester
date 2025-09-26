package protoexample

import (
	"fmt"
	"testing"
)

func TestNumber_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	x := FOO(1)
	if x.Number() != 1 {
		t.Errorf("x.Number() = %v, want %v", x.Number(), 1)
	}
}
