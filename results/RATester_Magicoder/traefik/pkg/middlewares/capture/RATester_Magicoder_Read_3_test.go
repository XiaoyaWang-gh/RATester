package capture

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

func TestRead_3(t *testing.T) {
	testCases := []struct {
		name    string
		source  io.ReadCloser
		content []byte
		want    int64
	}{
		{
			name:    "Read 1 byte",
			source:  ioutil.NopCloser(bytes.NewBuffer([]byte{'a'})),
			content: []byte{'a'},
			want:    1,
		},
		{
			name:    "Read 2 bytes",
			source:  ioutil.NopCloser(bytes.NewBuffer([]byte{'a', 'b'})),
			content: []byte{'a', 'b'},
			want:    2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rc := &readCounter{source: tc.source}
			n, err := rc.Read(tc.content)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if n != len(tc.content) {
				t.Errorf("got %d bytes, want %d", n, len(tc.content))
			}
			if rc.size != tc.want {
				t.Errorf("got size %d, want %d", rc.size, tc.want)
			}
		})
	}
}
