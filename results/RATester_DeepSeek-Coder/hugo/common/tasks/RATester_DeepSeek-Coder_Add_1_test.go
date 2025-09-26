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

	r := &RunEvery{
		HandleError: func(name string, err error) {
			t.Errorf("Function %s returned error: %v", name, err)
		},
	}

	f := Func{
		IntervalLow:  500 * time.Millisecond,
		IntervalHigh: 20 * time.Second,
		F: func(interval time.Duration) (time.Duration, error) {
			return interval, nil
		},
	}

	r.Add("test", f)

	if len(r.funcs) != 1 {
		t.Errorf("Expected 1 function, got %d", len(r.funcs))
	}

	if _, ok := r.funcs["test"]; !ok {
		t.Error("Expected function 'test' to be in the list")
	}

	if r.funcs["test"].interval != f.IntervalHigh/3 && r.funcs["test"].interval != f.IntervalLow {
		t.Errorf("Expected interval to be %v or %v, got %v", f.IntervalHigh/3, f.IntervalLow, r.funcs["test"].interval)
	}
}
