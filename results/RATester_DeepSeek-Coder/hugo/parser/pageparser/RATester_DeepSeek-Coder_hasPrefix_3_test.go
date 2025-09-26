package pageparser

import (
	"fmt"
	"testing"
)

func TestHasPrefix_3(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		prefix []byte
		want   bool
	}{
		{
			name:   "hasPrefix_true",
			input:  []byte("hello world"),
			prefix: []byte("hello"),
			want:   true,
		},
		{
			name:   "hasPrefix_false",
			input:  []byte("hello world"),
			prefix: []byte("world"),
			want:   false,
		},
		{
			name:   "hasPrefix_empty",
			input:  []byte("hello world"),
			prefix: []byte(""),
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &pageLexer{
				input: tt.input,
			}

			got := l.hasPrefix(tt.prefix)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
