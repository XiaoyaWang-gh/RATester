package dynacache

import (
	"fmt"
	"testing"
)

func TestWeightFraction_1(t *testing.T) {
	tests := []struct {
		name   string
		weight int
		want   float64
	}{
		{
			name:   "Weight is 50",
			weight: 50,
			want:   0.5,
		},
		{
			name:   "Weight is 100",
			weight: 100,
			want:   1.0,
		},
		{
			name:   "Weight is 0",
			weight: 0,
			want:   0.0,
		},
		{
			name:   "Weight is 1",
			weight: 1,
			want:   0.01,
		},
		{
			name:   "Weight is 1000",
			weight: 1000,
			want:   10.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			o := OptionsPartition{
				Weight: tt.weight,
			}
			if got := o.WeightFraction(); got != tt.want {
				t.Errorf("WeightFraction() = %v, want %v", got, tt.want)
			}
		})
	}
}
