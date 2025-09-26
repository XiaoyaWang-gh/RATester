package template

import (
	"fmt"
	"testing"
)

func TestisHTMLSpaceOrASCIIAlnum_1(t *testing.T) {
	tests := []struct {
		name  string
		input byte
		want  bool
	}{
		{"Test case 1", 'a', true},
		{"Test case 2", '1', true},
		{"Test case 3", ' ', true},
		{"Test case 4", '@', true},
		{"Test case 5", '~', false},
		{"Test case 6", 'é', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isHTMLSpaceOrASCIIAlnum(tt.input); got != tt.want {
				t.Errorf("isHTMLSpaceOrASCIIAlnum(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
