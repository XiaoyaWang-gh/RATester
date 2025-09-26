package captcha

import (
	"fmt"
	"testing"
)

func TestrandIntn_1(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Test case 1",
			n:    10,
			want: 5,
		},
		{
			name: "Test case 2",
			n:    20,
			want: 15,
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

			if got := randIntn(tt.n); got != tt.want {
				t.Errorf("randIntn() = %v, want %v", got, tt.want)
			}
		})
	}
}
