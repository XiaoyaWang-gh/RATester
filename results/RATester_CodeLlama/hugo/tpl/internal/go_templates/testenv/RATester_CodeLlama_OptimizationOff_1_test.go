package testenv

import (
	"fmt"
	"testing"
)

func TestOptimizationOff_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if OptimizationOff() != false {
		t.Errorf("OptimizationOff() = %v, want %v", OptimizationOff(), false)
	}
}
