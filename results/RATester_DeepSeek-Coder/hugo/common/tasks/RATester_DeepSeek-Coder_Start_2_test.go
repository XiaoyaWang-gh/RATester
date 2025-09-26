package tasks

import (
	"fmt"
	"testing"
	"time"
)

func TestStart_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &RunEvery{
		HandleError: func(name string, err error) {
			t.Errorf("Function %s returned an error: %v", name, err)
		},
		RunImmediately: true,
		funcs:          make(map[string]*Func),
	}

	r.Add("testFunc", Func{
		IntervalLow:  100 * time.Millisecond,
		IntervalHigh: 200 * time.Millisecond,
		F: func(interval time.Duration) (time.Duration, error) {
			return interval, nil
		},
	})

	if err := r.Start(); err != nil {
		t.Fatalf("Failed to start RunEvery: %v", err)
	}

	time.Sleep(300 * time.Millisecond)

	r.Close()

	// Add another function to check if it's correctly added and run
	r.Add("testFunc2", Func{
		IntervalLow:  300 * time.Millisecond,
		IntervalHigh: 400 * time.Millisecond,
		F: func(interval time.Duration) (time.Duration, error) {
			return interval, nil
		},
	})

	time.Sleep(500 * time.Millisecond)

	r.Close()
}
