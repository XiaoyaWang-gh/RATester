package tasks

import (
	"fmt"
	"testing"
	"time"
)

func TestAdd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &RunEvery{}

	// Test adding a function with default interval
	f1 := Func{
		F: func(interval time.Duration) (time.Duration, error) {
			return 0, nil
		},
	}
	r.Add("f1", f1)
	if _, ok := r.funcs["f1"]; !ok {
		t.Error("Expected function to be added")
	}
	if r.funcs["f1"].IntervalLow != 500*time.Millisecond {
		t.Error("Expected default interval low to be set")
	}
	if r.funcs["f1"].IntervalHigh != 20*time.Second {
		t.Error("Expected default interval high to be set")
	}

	// Test adding a function with custom interval
	f2 := Func{
		IntervalLow:  100 * time.Millisecond,
		IntervalHigh: 300 * time.Millisecond,
		F: func(interval time.Duration) (time.Duration, error) {
			return 0, nil
		},
	}
	r.Add("f2", f2)
	if _, ok := r.funcs["f2"]; !ok {
		t.Error("Expected function to be added")
	}
	if r.funcs["f2"].IntervalLow != 100*time.Millisecond {
		t.Error("Expected custom interval low to be set")
	}
	if r.funcs["f2"].IntervalHigh != 300*time.Millisecond {
		t.Error("Expected custom interval high to be set")
	}
}
