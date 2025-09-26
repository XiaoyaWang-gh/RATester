package pageparser

import (
	"fmt"
	"testing"
)

func TestindexNonWhiteSpace_1(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		in   rune
		want int
	}{
		{
			name: "Test case 1",
			s:    []byte("Hello World"),
			in:   'W',
			want: 6,
		},
		{
			name: "Test case 2",
			s:    []byte("Hello World"),
			in:   'Z',
			want: -1,
		},
		{
			name: "Test case 3",
			s:    []byte("Hello World"),
			in:   ' ',
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := indexNonWhiteSpace(tt.s, tt.in); got != tt.want {
				t.Errorf("indexNonWhiteSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
