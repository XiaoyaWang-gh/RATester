package template

import (
	"fmt"
	"testing"
)

func TestescFnsEq_1(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{
			name: "Test case 1",
			a:    "a",
			b:    "b",
			want: false,
		},
		{
			name: "Test case 2",
			a:    "a",
			b:    "a",
			want: true,
		},
		{
			name: "Test case 3",
			a:    "a",
			b:    "A",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := escFnsEq(tt.a, tt.b); got != tt.want {
				t.Errorf("escFnsEq() = %v, want %v", got, tt.want)
			}
		})
	}
}
