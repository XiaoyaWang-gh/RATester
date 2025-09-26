package context

import (
	"fmt"
	"testing"
)

func TestIsCachable_1(t *testing.T) {
	tests := []struct {
		name   string
		output *BeegoOutput
		want   bool
	}{
		{
			name: "Test case 1",
			output: &BeegoOutput{
				Status: 200,
			},
			want: true,
		},
		{
			name: "Test case 2",
			output: &BeegoOutput{
				Status: 304,
			},
			want: true,
		},
		{
			name: "Test case 3",
			output: &BeegoOutput{
				Status: 404,
			},
			want: false,
		},
		{
			name: "Test case 4",
			output: &BeegoOutput{
				Status: 500,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.output.IsCachable(); got != tt.want {
				t.Errorf("IsCachable() = %v, want %v", got, tt.want)
			}
		})
	}
}
