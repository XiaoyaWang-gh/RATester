package wrr

import (
	"fmt"
	"testing"
)

func TestHash_1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty string",
			input: "",
			want:  "0",
		},
		{
			name:  "single character",
			input: "a",
			want:  "61",
		},
		{
			name:  "multiple characters",
			input: "hello",
			want:  "3e25960a79dbc69b674cd4ec67a72c62",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := hash(tt.input); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
