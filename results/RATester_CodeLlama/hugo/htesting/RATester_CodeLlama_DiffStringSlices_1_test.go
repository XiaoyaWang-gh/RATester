package htesting

import (
	"fmt"
	"testing"
)

func TestDiffStringSlices_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	slice1 := []string{"a", "b", "c", "d", "e"}
	slice2 := []string{"a", "b", "c", "d", "e"}
	diffStr := DiffStringSlices(slice1, slice2)
	if len(diffStr) != 0 {
		t.Errorf("DiffStringSlices() = %v, want %v", diffStr, 0)
	}
}
