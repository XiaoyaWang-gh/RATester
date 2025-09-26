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

	want := false
	got := OptimizationOff()

	if got != want {
		t.Errorf("OptimizationOff() = %v, want %v", got, want)
	}
}
