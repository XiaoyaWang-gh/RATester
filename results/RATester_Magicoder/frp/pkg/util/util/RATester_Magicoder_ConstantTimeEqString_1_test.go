package util

import (
	"fmt"
	"testing"
)

func TestConstantTimeEqString_1(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{
			name: "Equal strings",
			a:    "hello",
			b:    "hello",
			want: true,
		},
		{
			name: "Unequal strings",
			a:    "hello",
			b:    "world",
			want: false,
		},
		{
			name: "Empty strings",
			a:    "",
			b:    "",
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

			if got := ConstantTimeEqString(tt.a, tt.b); got != tt.want {
				t.Errorf("ConstantTimeEqString() = %v, want %v", got, tt.want)
			}
		})
	}
}
