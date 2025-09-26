package logs

import (
	"fmt"
	"testing"
)

func TestneedRotateDaily_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		MaxLines:         1000,
		maxLinesCurLines: 1001,
		Daily:            true,
		dailyOpenDate:    1,
	}

	tests := []struct {
		day  int
		want bool
	}{
		{1, false},
		{2, true},
	}

	for _, test := range tests {
		if got := w.needRotateDaily(test.day); got != test.want {
			t.Errorf("needRotateDaily(%v) = %v", test.day, got)
		}
	}
}
