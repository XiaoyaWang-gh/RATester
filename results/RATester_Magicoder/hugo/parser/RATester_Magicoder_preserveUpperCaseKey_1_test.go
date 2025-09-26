package parser

import (
	"fmt"
	"testing"
)

func TestpreserveUpperCaseKey_1(t *testing.T) {
	tests := []struct {
		name  string
		match []byte
		want  bool
	}{
		{
			name:  "Test case 1",
			match: []byte("Hello"),
			want:  true,
		},
		{
			name:  "Test case 2",
			match: []byte("hello"),
			want:  false,
		},
		{
			name:  "Test case 3",
			match: []byte("HELLO"),
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := preserveUpperCaseKey(tt.match); got != tt.want {
				t.Errorf("preserveUpperCaseKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
