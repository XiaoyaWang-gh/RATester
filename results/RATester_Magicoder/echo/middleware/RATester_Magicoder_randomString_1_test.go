package middleware

import (
	"fmt"
	"testing"
)

func TestRandomString_1(t *testing.T) {
	tests := []struct {
		name   string
		length uint8
		want   string
	}{
		{
			name:   "Test case 1",
			length: 10,
			want:   "random string",
		},
		{
			name:   "Test case 2",
			length: 20,
			want:   "another random string",
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

			if got := randomString(tt.length); got != tt.want {
				t.Errorf("randomString() = %v, want %v", got, tt.want)
			}
		})
	}
}
