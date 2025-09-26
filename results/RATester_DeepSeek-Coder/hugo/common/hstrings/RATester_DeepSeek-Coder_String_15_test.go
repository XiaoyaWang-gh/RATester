package hstrings

import (
	"fmt"
	"testing"
)

func TestString_15(t *testing.T) {
	tests := []struct {
		name string
		s    StringEqualFold
		want string
	}{
		{
			name: "Test Case 1",
			s:    "Hello",
			want: "Hello",
		},
		{
			name: "Test Case 2",
			s:    "WORLD",
			want: "WORLD",
		},
		{
			name: "Test Case 3",
			s:    "1234567890",
			want: "1234567890",
		},
		{
			name: "Test Case 4",
			s:    "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
