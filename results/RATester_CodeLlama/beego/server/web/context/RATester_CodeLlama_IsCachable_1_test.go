package context

import (
	"fmt"
	"testing"
)

func TestIsCachable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	output := &BeegoOutput{}
	output.Status = 200
	if !output.IsCachable() {
		t.Errorf("IsCachable() = %v, want %v", output.IsCachable(), true)
	}

	output.Status = 201
	if output.IsCachable() {
		t.Errorf("IsCachable() = %v, want %v", output.IsCachable(), false)
	}

	output.Status = 304
	if !output.IsCachable() {
		t.Errorf("IsCachable() = %v, want %v", output.IsCachable(), true)
	}
}
