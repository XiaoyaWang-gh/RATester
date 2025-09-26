package template

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDecodeCSS_1(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{
			name: "test case 1",
			s:    []byte(`\\48\\65\\6c\\6c\\6f`),
			want: []byte("Hello"),
		},
		{
			name: "test case 2",
			s:    []byte(`\\77\\6f\\72\\6c\\64`),
			want: []byte("world"),
		},
		{
			name: "test case 3",
			s:    []byte(`\\5c\\5c`),
			want: []byte("\\\\"),
		},
		{
			name: "test case 4",
			s:    []byte(`\\5c\\5c\\5c\\5c`),
			want: []byte("\\\\\\\\"),
		},
		{
			name: "test case 5",
			s:    []byte(`\\5c\\5c\\5c\\5c\\5c`),
			want: []byte("\\\\\\\\\\\\"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := decodeCSS(tt.s); !bytes.Equal(got, tt.want) {
				t.Errorf("decodeCSS() = %v, want %v", got, tt.want)
			}
		})
	}
}
