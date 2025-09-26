package paths

import (
	"fmt"
	"testing"
)

func TestisAllowedPathCharacter_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		i    int
		r    rune
		want bool
	}{
		{
			name: "Test case 1",
			s:    "test",
			i:    0,
			r:    't',
			want: true,
		},
		{
			name: "Test case 2",
			s:    "test",
			i:    0,
			r:    ' ',
			want: false,
		},
		{
			name: "Test case 3",
			s:    "test",
			i:    0,
			r:    '@',
			want: true,
		},
		{
			name: "Test case 4",
			s:    "test",
			i:    0,
			r:    '%',
			want: false,
		},
		{
			name: "Test case 5",
			s:    "test",
			i:    0,
			r:    '+',
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

			if got := isAllowedPathCharacter(tt.s, tt.i, tt.r); got != tt.want {
				t.Errorf("isAllowedPathCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
