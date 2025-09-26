package metrics

import (
	"fmt"
	"testing"
	"time"
)

func TestMeasureSince_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Store{}
	key := "key"
	start := time.Now()
	s.MeasureSince(key, start)
	if len(s.metrics[key]) != 1 {
		t.Errorf("expected 1 measurement, got %d", len(s.metrics[key]))
	}
}
