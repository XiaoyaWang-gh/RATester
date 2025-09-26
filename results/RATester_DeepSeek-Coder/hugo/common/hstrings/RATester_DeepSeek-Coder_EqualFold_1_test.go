package hstrings

import (
	"fmt"
	"testing"
)

func TestEqualFold_1(t *testing.T) {
	tests := []struct {
		name string
		s    StringEqualFold
		arg  string
		want bool
	}{
		{
			name: "equal strings",
			s:    "Hello",
			arg:  "hello",
			want: true,
		},
		{
			name: "unequal strings",
			s:    "Hello",
			arg:  "world",
			want: false,
		},
		{
			name: "empty strings",
			s:    "",
			arg:  "",
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

			if got := tt.s.EqualFold(tt.arg); got != tt.want {
				t.Errorf("EqualFold() = %v, want %v", got, tt.want)
			}
		})
	}
}
