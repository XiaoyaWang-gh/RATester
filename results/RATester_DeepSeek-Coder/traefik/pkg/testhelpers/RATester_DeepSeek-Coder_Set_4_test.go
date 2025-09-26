package testhelpers

import (
	"fmt"
	"testing"
)

func TestSet_4(t *testing.T) {
	tests := []struct {
		name   string
		value  float64
		expect float64
	}{
		{
			name:   "set positive value",
			value:  1.23,
			expect: 1.23,
		},
		{
			name:   "set negative value",
			value:  -4.56,
			expect: -4.56,
		},
		{
			name:   "set zero value",
			value:  0,
			expect: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gauge := &CollectingGauge{}
			gauge.Set(test.value)

			if gauge.GaugeValue != test.expect {
				t.Errorf("expected %v, got %v", test.expect, gauge.GaugeValue)
			}
		})
	}
}
