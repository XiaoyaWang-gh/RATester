package csrf

import (
	"fmt"
	"testing"
)

func TestcompareTokens_1(t *testing.T) {
	tests := []struct {
		name string
		a    []byte
		b    []byte
		want bool
	}{
		{
			name: "Test case 1",
			a:    []byte("token1"),
			b:    []byte("token1"),
			want: true,
		},
		{
			name: "Test case 2",
			a:    []byte("token2"),
			b:    []byte("token3"),
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

			if got := compareTokens(tt.a, tt.b); got != tt.want {
				t.Errorf("compareTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}
