package tasks

import (
	"fmt"
	"testing"
	"time"
)

func Testrun_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &RunEvery{
		HandleError: func(name string, err error) {
			t.Errorf("Error in function %s: %v", name, err)
		},
		funcs: map[string]*Func{
			"testFunc": {
				IntervalLow:  time.Second,
				IntervalHigh: time.Minute,
				F:            func(interval time.Duration) (time.Duration, error) { return interval, nil },
				interval:     time.Second,
				last:         time.Now().Add(-2 * time.Second),
			},
		},
	}

	r.run()

	// Add more assertions or tests as needed
}
