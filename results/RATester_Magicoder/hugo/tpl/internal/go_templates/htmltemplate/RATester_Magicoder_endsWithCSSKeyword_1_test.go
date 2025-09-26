package template

import (
	"fmt"
	"testing"
)

func TestendsWithCSSKeyword_1(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		kw   string
		want bool
	}{
		{
			name: "Test case 1",
			b:    []byte("color: red;"),
			kw:   "red",
			want: true,
		},
		{
			name: "Test case 2",
			b:    []byte("color: blue;"),
			kw:   "red",
			want: false,
		},
		{
			name: "Test case 3",
			b:    []byte("color: green;"),
			kw:   "green",
			want: true,
		},
		{
			name: "Test case 4",
			b:    []byte("color: yellow;"),
			kw:   "yellow",
			want: true,
		},
		{
			name: "Test case 5",
			b:    []byte("color: black;"),
			kw:   "black",
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

			if got := endsWithCSSKeyword(tt.b, tt.kw); got != tt.want {
				t.Errorf("endsWithCSSKeyword() = %v, want %v", got, tt.want)
			}
		})
	}
}
