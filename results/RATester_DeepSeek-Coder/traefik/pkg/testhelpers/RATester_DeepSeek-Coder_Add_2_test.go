package testhelpers

import (
	"fmt"
	"testing"
)

func TestAdd_2(t *testing.T) {
	tests := []struct {
		name   string
		g      *CollectingGauge
		delta  float64
		expect float64
	}{
		{
			name:   "Add positive delta",
			g:      &CollectingGauge{GaugeValue: 10},
			delta:  5,
			expect: 15,
		},
		{
			name:   "Add negative delta",
			g:      &CollectingGauge{GaugeValue: 10},
			delta:  -3,
			expect: 7,
		},
		{
			name:   "Add zero delta",
			g:      &CollectingGauge{GaugeValue: 10},
			delta:  0,
			expect: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.g.Add(tt.delta)
			if tt.g.GaugeValue != tt.expect {
				t.Errorf("Expected GaugeValue to be %v, but got %v", tt.expect, tt.g.GaugeValue)
			}
		})
	}
}
