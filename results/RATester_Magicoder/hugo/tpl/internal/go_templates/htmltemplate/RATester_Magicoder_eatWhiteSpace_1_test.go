package template

import (
	"fmt"
	"testing"
)

func TesteatWhiteSpace_1(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		i    int
		want int
	}{
		{
			name: "Test case 1",
			s:    []byte("Hello World"),
			i:    0,
			want: 5,
		},
		{
			name: "Test case 2",
			s:    []byte("  Hello World"),
			i:    0,
			want: 1,
		},
		{
			name: "Test case 3",
			s:    []byte("Hello\tWorld"),
			i:    0,
			want: 5,
		},
		{
			name: "Test case 4",
			s:    []byte("Hello\nWorld"),
			i:    0,
			want: 5,
		},
		{
			name: "Test case 5",
			s:    []byte("Hello\fWorld"),
			i:    0,
			want: 5,
		},
		{
			name: "Test case 6",
			s:    []byte("Hello\rWorld"),
			i:    0,
			want: 5,
		},
		{
			name: "Test case 7",
			s:    []byte("Hello World"),
			i:    5,
			want: 11,
		},
		{
			name: "Test case 8",
			s:    []byte("Hello World"),
			i:    11,
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := eatWhiteSpace(tt.s, tt.i); got != tt.want {
				t.Errorf("eatWhiteSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
