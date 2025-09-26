package validation

import (
	"fmt"
	"testing"
)

func TestNumIn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "Test"
	num, err := numIn(name)
	if err != nil {
		t.Errorf("numIn() error = %v", err)
		return
	}
	if num != 1 {
		t.Errorf("numIn() = %v, want %v", num, 1)
	}
}
