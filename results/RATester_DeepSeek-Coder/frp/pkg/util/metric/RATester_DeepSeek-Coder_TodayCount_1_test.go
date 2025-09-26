package metric

import (
	"fmt"
	"testing"
	"time"
)

func TestTodayCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	counter := &StandardDateCounter{
		reserveDays:    7,
		counts:         make([]int64, 7),
		lastUpdateDate: time.Now(),
	}

	counter.Inc(10)
	expected := counter.counts[0]
	actual := counter.TodayCount()
	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
