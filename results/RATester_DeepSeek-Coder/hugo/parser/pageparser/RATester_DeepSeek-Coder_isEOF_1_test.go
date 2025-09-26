package pageparser

import (
	"fmt"
	"testing"
)

func TestIsEOF_1(t *testing.T) {
	tests := []struct {
		name string
		l    *pageLexer
		want bool
	}{
		{
			name: "Test case 1",
			l: &pageLexer{
				input: []byte("Hello, world!"),
				pos:   0,
			},
			want: false,
		},
		{
			name: "Test case 2",
			l: &pageLexer{
				input: []byte("Hello, world!"),
				pos:   13,
			},
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

			if got := tt.l.isEOF(); got != tt.want {
				t.Errorf("pageLexer.isEOF() = %v, want %v", got, tt.want)
			}
		})
	}
}
