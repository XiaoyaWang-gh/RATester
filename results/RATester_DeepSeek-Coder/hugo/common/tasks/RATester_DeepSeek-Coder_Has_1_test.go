package tasks

import (
	"fmt"
	"testing"
	"time"
)

func TestHas_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &RunEvery{
		funcs: make(map[string]*Func),
	}

	// Add a function
	r.Add("testFunc", Func{
		IntervalLow:  1 * time.Second,
		IntervalHigh: 2 * time.Second,
		F: func(interval time.Duration) (time.Duration, error) {
			return interval, nil
		},
	})

	// Test Has method
	has := r.Has("testFunc")
	if !has {
		t.Errorf("Expected Has to return true for existing function")
	}

	// Test Has method for non-existing function
	has = r.Has("nonExistingFunc")
	if has {
		t.Errorf("Expected Has to return false for non-existing function")
	}
}
