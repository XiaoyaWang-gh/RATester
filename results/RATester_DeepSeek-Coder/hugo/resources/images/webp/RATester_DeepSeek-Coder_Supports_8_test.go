package webp

import (
	"fmt"
	"testing"
)

func TestSupports_8(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Test case 1",
			want: false,
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Supports(); got != tt.want {
				t.Errorf("Supports() = %v, want %v", got, tt.want)
			}
		})
	}
}
