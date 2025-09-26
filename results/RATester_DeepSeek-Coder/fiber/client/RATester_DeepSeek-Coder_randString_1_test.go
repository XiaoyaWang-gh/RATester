package client

import (
	"fmt"
	"testing"
)

func TestRandString_1(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "Test case 1",
			n:    10,
			want: "abcdefghij",
		},
		{
			name: "Test case 2",
			n:    20,
			want: "abcdefghijklmnopqrst",
		},
		{
			name: "Test case 3",
			n:    30,
			want: "abcdefghijklmnopqrstuvwxyz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := randString(tt.n); got != tt.want {
				t.Errorf("randString() = %v, want %v", got, tt.want)
			}
		})
	}
}
