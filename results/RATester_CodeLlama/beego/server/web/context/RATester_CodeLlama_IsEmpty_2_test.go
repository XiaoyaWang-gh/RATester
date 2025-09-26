package context

import (
	"fmt"
	"testing"
)

func TestIsEmpty_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{}
	output.Status = 201
	if !output.IsEmpty() {
		t.Errorf("TestIsEmpty failed")
	}
	output.Status = 204
	if !output.IsEmpty() {
		t.Errorf("TestIsEmpty failed")
	}
	output.Status = 304
	if !output.IsEmpty() {
		t.Errorf("TestIsEmpty failed")
	}
}
