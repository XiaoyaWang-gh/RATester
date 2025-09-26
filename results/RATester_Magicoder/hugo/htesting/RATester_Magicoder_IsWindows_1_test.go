package htesting

import (
	"fmt"
	"testing"
)

func TestIsWindows_1(t *testing.T) {
	type test struct {
		name string
		want bool
	}

	tests := []test{
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

			if got := IsWindows(); got != tt.want {
				t.Errorf("IsWindows() = %v, want %v", got, tt.want)
			}
		})
	}
}
