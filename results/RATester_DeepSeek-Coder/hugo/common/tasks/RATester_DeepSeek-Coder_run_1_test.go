package tasks

import (
	"fmt"
	"testing"
	"time"
)

func TestRun_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &RunEvery{
		HandleError: func(name string, err error) {
			t.Errorf("Function %s returned an error: %v", name, err)
		},
		funcs: map[string]*Func{
			"test": {
				IntervalLow:  1 * time.Second,
				IntervalHigh: 10 * time.Second,
				F: func(interval time.Duration) (time.Duration, error) {
					return 5 * time.Second, nil
				},
			},
		},
		quit: make(chan struct{}),
	}

	r.run()

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.funcs["test"].last.IsZero() {
		t.Error("Expected last run time to be set")
	}

	if r.funcs["test"].interval != 5*time.Second {
		t.Errorf("Expected interval to be 5s, got %v", r.funcs["test"].interval)
	}
}
