package protoexample

import (
	"fmt"
	"testing"
)

func TestString_2(t *testing.T) {
	tests := []struct {
		name string
		x    FOO
		want string
	}{
		{
			name: "Test Case 1",
			x:    FOO(1),
			want: "Test Case 1",
		},
		{
			name: "Test Case 2",
			x:    FOO(2),
			want: "Test Case 2",
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

			if got := tt.x.String(); got != tt.want {
				t.Errorf("FOO.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
