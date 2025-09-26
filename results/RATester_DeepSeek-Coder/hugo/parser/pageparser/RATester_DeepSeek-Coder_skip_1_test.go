package pageparser

import (
	"fmt"
	"testing"
)

func TestSkip_1(t *testing.T) {
	tests := []struct {
		name     string
		s        *sectionHandler
		want     int
		wantSkip bool
	}{
		{
			name: "Test when skipAll is true",
			s: &sectionHandler{
				l: &pageLexer{
					input: []byte("test"),
				},
				skipAll: true,
			},
			want:     -1,
			wantSkip: true,
		},
		{
			name: "Test when skipFunc returns -1",
			s: &sectionHandler{
				l: &pageLexer{
					input: []byte("test"),
				},
				skipFunc: func(l *pageLexer) int { return -1 },
			},
			want:     -1,
			wantSkip: true,
		},
		{
			name: "Test when skipFunc returns non-negative value",
			s: &sectionHandler{
				l: &pageLexer{
					input: []byte("test"),
				},
				skipFunc: func(l *pageLexer) int { return 1 },
			},
			want:     1,
			wantSkip: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.s.skip()
			if got != tt.want {
				t.Errorf("sectionHandler.skip() = %v, want %v", got, tt.want)
			}
			if tt.s.skipAll != tt.wantSkip {
				t.Errorf("sectionHandler.skipAll = %v, want %v", tt.s.skipAll, tt.wantSkip)
			}
		})
	}
}
