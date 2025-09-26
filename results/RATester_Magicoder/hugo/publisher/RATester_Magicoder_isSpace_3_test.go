package publisher

import (
	"fmt"
	"testing"
)

func TestisSpace_3(t *testing.T) {
	tests := []struct {
		name  string
		input byte
		want  bool
	}{
		{"space", ' ', true},
		{"tab", '\t', true},
		{"newline", '\n', true},
		{"a", 'a', false},
		{"0", '0', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isSpace(tt.input); got != tt.want {
				t.Errorf("isSpace(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
