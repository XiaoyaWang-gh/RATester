package urlreplacers

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestcheckCandidateBase_1(t *testing.T) {
	tests := []struct {
		name string
		l    *absurllexer
		want *absurllexer
	}{
		{
			name: "test case 1",
			l: &absurllexer{
				content: []byte("http://example.com/path"),
				w:       bytes.NewBuffer([]byte{}),
				path:    []byte("."),
				pos:     0,
				start:   0,
			},
			want: &absurllexer{
				content: []byte("http://example.com/path"),
				w:       bytes.NewBuffer([]byte{}),
				path:    []byte("."),
				pos:     0,
				start:   0,
			},
		},
		{
			name: "test case 2",
			l: &absurllexer{
				content: []byte("http://example.com/path"),
				w:       bytes.NewBuffer([]byte{}),
				path:    []byte("."),
				pos:     0,
				start:   0,
			},
			want: &absurllexer{
				content: []byte("http://example.com/path"),
				w:       bytes.NewBuffer([]byte{}),
				path:    []byte("."),
				pos:     0,
				start:   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			checkCandidateBase(tt.l)
			if !reflect.DeepEqual(tt.l, tt.want) {
				t.Errorf("checkCandidateBase() = %v, want %v", tt.l, tt.want)
			}
		})
	}
}
