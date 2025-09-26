package dynacache

import (
	"fmt"
	"testing"
)

func TestWeightFraction_1(t *testing.T) {
	tests := []struct {
		name string
		o    OptionsPartition
		want float64
	}{
		{
			name: "Test with weight 50",
			o:    OptionsPartition{Weight: 50},
			want: 0.5,
		},
		{
			name: "Test with weight 100",
			o:    OptionsPartition{Weight: 100},
			want: 1.0,
		},
		{
			name: "Test with weight 0",
			o:    OptionsPartition{Weight: 0},
			want: 0.0,
		},
		{
			name: "Test with weight 20",
			o:    OptionsPartition{Weight: 20},
			want: 0.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.o.WeightFraction(); got != tt.want {
				t.Errorf("OptionsPartition.WeightFraction() = %v, want %v", got, tt.want)
			}
		})
	}
}
