package util

import (
	"fmt"
	"testing"
	"time"
)

func TestRandomSleep_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	duration := time.Second
	minRatio := 0.5
	maxRatio := 1.0
	d := RandomSleep(duration, minRatio, maxRatio)
	if d < 0 {
		t.Errorf("RandomSleep(%v, %v, %v) = %v < 0", duration, minRatio, maxRatio, d)
	}
	if d > duration {
		t.Errorf("RandomSleep(%v, %v, %v) = %v > %v", duration, minRatio, maxRatio, d, duration)
	}
	if d.Nanoseconds()%1000 != 0 {
		t.Errorf("RandomSleep(%v, %v, %v) = %v", duration, minRatio, maxRatio, d)
	}
}
