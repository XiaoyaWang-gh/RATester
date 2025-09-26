package htesting

import (
	"fmt"
	"testing"
)

func TestIsCI_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Test case 1",
			want: false,
		},
		{
			name: "Test case 2",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsCI(); got != tt.want {
				t.Errorf("IsCI() = %v, want %v", got, tt.want)
			}
		})
	}
}
