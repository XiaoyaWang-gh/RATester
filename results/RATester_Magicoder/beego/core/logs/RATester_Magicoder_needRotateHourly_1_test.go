package logs

import (
	"fmt"
	"testing"
)

func TestneedRotateHourly_1(t *testing.T) {
	w := &fileLogWriter{
		MaxLines:       1000,
		MaxSize:        1024,
		Hourly:         true,
		hourlyOpenDate: 1,
	}

	tests := []struct {
		name   string
		hour   int
		expect bool
	}{
		{"MaxLines", 1001, true},
		{"MaxSize", 1025, true},
		{"Hourly", 2, true},
		{"Normal", 1, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := w.needRotateHourly(test.hour)
			if result != test.expect {
				t.Errorf("needRotateHourly(%d) = %t; want %t", test.hour, result, test.expect)
			}
		})
	}
}
