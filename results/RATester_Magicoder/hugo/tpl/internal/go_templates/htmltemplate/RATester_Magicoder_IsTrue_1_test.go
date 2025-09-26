package template

import (
	"fmt"
	"testing"
)

func TestIsTrue_1(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  bool
	}{
		{
			name:  "Test case 1",
			input: 1,
			want:  true,
		},
		{
			name:  "Test case 2",
			input: "hello",
			want:  true,
		},
		{
			name:  "Test case 3",
			input: false,
			want:  false,
		},
		{
			name:  "Test case 4",
			input: nil,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, _ := IsTrue(tt.input)
			if got != tt.want {
				t.Errorf("IsTrue(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
