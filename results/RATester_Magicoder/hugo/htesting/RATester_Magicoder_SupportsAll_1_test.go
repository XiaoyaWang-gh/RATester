package htesting

import (
	"fmt"
	"testing"
)

func TestSupportsAll_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Test case 1",
			want: true,
		},
		{
			name: "Test case 2",
			want: false,
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

			if got := SupportsAll(); got != tt.want {
				t.Errorf("SupportsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
