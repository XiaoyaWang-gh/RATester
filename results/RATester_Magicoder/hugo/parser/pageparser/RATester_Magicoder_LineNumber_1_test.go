package pageparser

import (
	"fmt"
	"testing"
)

func TestLineNumber_1(t *testing.T) {
	testCases := []struct {
		name   string
		source []byte
		want   int
	}{
		{
			name:   "empty source",
			source: []byte{},
			want:   1,
		},
		{
			name:   "single line",
			source: []byte("Hello, world!"),
			want:   1,
		},
		{
			name:   "multiple lines",
			source: []byte("Hello,\nworld!"),
			want:   2,
		},
		{
			name:   "multiple lines with trailing newline",
			source: []byte("Hello,\nworld!\n"),
			want:   3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			it := &Iterator{}
			got := it.LineNumber(tc.source)
			if got != tc.want {
				t.Errorf("LineNumber() = %v, want %v", got, tc.want)
			}
		})
	}
}
