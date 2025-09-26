package logs

import (
	"fmt"
	"testing"
)

func TestNeedRotateHourly_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		MaxLines:         1000,
		maxLinesCurLines: 1001,
		Hourly:           true,
		hourlyOpenDate:   10,
	}

	hour := 11
	if !w.needRotateHourly(hour) {
		t.Errorf("Expected true, got false")
	}

	w.hourlyOpenDate = 11
	if w.needRotateHourly(hour) {
		t.Errorf("Expected false, got true")
	}
}
