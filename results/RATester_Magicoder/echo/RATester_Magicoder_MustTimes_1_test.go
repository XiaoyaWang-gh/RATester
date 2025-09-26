package echo

import (
	"fmt"
	"testing"
	"time"
)

func TestMustTimes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "2006-01-02T15:04:05Z"
		},
	}

	var times []time.Time
	b.MustTimes("test", &times, time.RFC3339)

	if len(times) != 1 {
		t.Errorf("Expected 1 time, got %d", len(times))
	}

	expectedTime, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if !times[0].Equal(expectedTime) {
		t.Errorf("Expected time %v, got %v", expectedTime, times[0])
	}
}
