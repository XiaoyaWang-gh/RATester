package math

import (
	"fmt"
	"testing"
)

func TestAbs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var (
		ns = &Namespace{}
		n  = 1.0
	)

	result, err := ns.Abs(n)
	if err != nil {
		t.Fatal(err)
	}

	if result != 1.0 {
		t.Errorf("Expected 1.0, got %v", result)
	}
}
