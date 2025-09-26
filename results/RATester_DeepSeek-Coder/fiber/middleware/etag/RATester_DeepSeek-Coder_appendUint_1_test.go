package etag

import (
	"bytes"
	"fmt"
	"testing"
)

func TestAppendUint_1(t *testing.T) {
	tests := []struct {
		name string
		dst  []byte
		n    uint32
		want []byte
	}{
		{
			name: "Test case 1",
			dst:  []byte("Hello, "),
			n:    123,
			want: []byte("Hello, 123"),
		},
		{
			name: "Test case 2",
			dst:  []byte("World, "),
			n:    456,
			want: []byte("World, 456"),
		},
		{
			name: "Test case 3",
			dst:  []byte("Hello, "),
			n:    789,
			want: []byte("Hello, 789"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := appendUint(tt.dst, tt.n)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("appendUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
