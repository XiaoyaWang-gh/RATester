package publisher

import (
	"fmt"
	"testing"
)

func Testnext_6(t *testing.T) {
	tests := []struct {
		name string
		l    *htmlElementsCollectorWriter
		want rune
	}{
		{
			name: "Test case 1",
			l: &htmlElementsCollectorWriter{
				input: []byte("Hello, World!"),
			},
			want: 'H',
		},
		{
			name: "Test case 2",
			l: &htmlElementsCollectorWriter{
				input: []byte(""),
			},
			want: eof,
		},
		{
			name: "Test case 3",
			l: &htmlElementsCollectorWriter{
				input: []byte("😀"),
			},
			want: '😀',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.l.next(); got != tt.want {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}
