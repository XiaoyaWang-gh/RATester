package metric

import (
	"fmt"
	"testing"
)

func TestCount_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want int32
	}{
		{
			name: "Test case 1",
			want: 10,
		},
		{
			name: "Test case 2",
			want: 20,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &StandardCounter{
				count: tt.want,
			}
			if got := c.Count(); got != tt.want {
				t.Errorf("StandardCounter.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
