package protoexample

import (
	"fmt"
	"testing"
)

func TestReset_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var x Test_OptionalGroup
	x.Reset()
	if x.RequiredField != nil {
		t.Errorf("x.RequiredField = %v, want nil", x.RequiredField)
	}
}
