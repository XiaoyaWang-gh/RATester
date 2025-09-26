package fiber

import (
	"fmt"
	"testing"
)

func TestIndexRune_1(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		needle int32
		want   bool
	}{
		{
			name:   "Test case 1",
			str:    "hello",
			needle: 'h',
			want:   true,
		},
		{
			name:   "Test case 2",
			str:    "world",
			needle: 'w',
			want:   true,
		},
		{
			name:   "Test case 3",
			str:    "test",
			needle: 'z',
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IndexRune(tt.str, tt.needle); got != tt.want {
				t.Errorf("IndexRune() = %v, want %v", got, tt.want)
			}
		})
	}
}
