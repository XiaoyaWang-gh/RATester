package urlreplacers

import (
	"fmt"
	"testing"
)

func TestposAfterURL_1(t *testing.T) {
	l := &absurllexer{
		content: []byte("test content"),
		pos:     0,
	}

	tests := []struct {
		name string
		q    []byte
		want int
	}{
		{
			name: "empty query",
			q:    []byte(""),
			want: 0,
		},
		{
			name: "non-empty query",
			q:    []byte("content"),
			want: 10,
		},
		{
			name: "query not found",
			q:    []byte("not found"),
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

			if got := l.posAfterURL(tt.q); got != tt.want {
				t.Errorf("posAfterURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
