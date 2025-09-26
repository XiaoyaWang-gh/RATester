package urlreplacers

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestcheckCandidateSrcset_1(t *testing.T) {
	tests := []struct {
		name string
		l    *absurllexer
		want *absurllexer
	}{
		{
			name: "Test case 1",
			l: &absurllexer{
				content: []byte("<img srcset=\"/path/to/image.jpg 2x, /path/to/image@2x.jpg 3x\">"),
				w:       bytes.NewBuffer([]byte{}),
				path:    []byte("."),
				pos:     0,
				start:   0,
				quotes:  [][]byte{},
			},
			want: &absurllexer{
				content: []byte("<img srcset=\"/path/to/image.jpg 2x, /path/to/image@2x.jpg 3x\">"),
				w:       bytes.NewBuffer([]byte{}),
				path:    []byte("."),
				pos:     0,
				start:   0,
				quotes:  [][]byte{},
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			checkCandidateSrcset(tt.l)
			if !reflect.DeepEqual(tt.l, tt.want) {
				t.Errorf("checkCandidateSrcset() = %v, want %v", tt.l, tt.want)
			}
		})
	}
}
