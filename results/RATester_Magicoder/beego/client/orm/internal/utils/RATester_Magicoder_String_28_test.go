package utils

import (
	"fmt"
	"testing"
)

func TestString_28(t *testing.T) {
	tests := []struct {
		name string
		f    StrTo
		want string
	}{
		{
			name: "Test case 1",
			f:    "test",
			want: "test",
		},
		{
			name: "Test case 2",
			f:    "",
			want: "",
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

			if got := tt.f.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
