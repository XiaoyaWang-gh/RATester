package htesting

import (
	"fmt"
	"testing"
)

func TestRandBool_1(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := RandBool()
			if got != tt.want {
				t.Errorf("RandBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
