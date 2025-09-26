package context

import (
	"fmt"
	"testing"
)

func TestIsSuccessful_1(t *testing.T) {
	tests := []struct {
		name   string
		output *BeegoOutput
		want   bool
	}{
		{
			name:   "Test case 1",
			output: &BeegoOutput{Status: 200},
			want:   true,
		},
		{
			name:   "Test case 2",
			output: &BeegoOutput{Status: 201},
			want:   true,
		},
		{
			name:   "Test case 3",
			output: &BeegoOutput{Status: 199},
			want:   false,
		},
		{
			name:   "Test case 4",
			output: &BeegoOutput{Status: 300},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.output.IsSuccessful(); got != tt.want {
				t.Errorf("BeegoOutput.IsSuccessful() = %v, want %v", got, tt.want)
			}
		})
	}
}
