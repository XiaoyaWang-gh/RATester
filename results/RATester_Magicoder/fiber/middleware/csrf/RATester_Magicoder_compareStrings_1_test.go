package csrf

import (
	"fmt"
	"testing"
)

func TestcompareStrings_1(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{
			name: "equal strings",
			a:    "hello",
			b:    "hello",
			want: true,
		},
		{
			name: "unequal strings",
			a:    "hello",
			b:    "world",
			want: false,
		},
		{
			name: "empty strings",
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

			if got := compareStrings(tt.a, tt.b); got != tt.want {
				t.Errorf("compareStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
