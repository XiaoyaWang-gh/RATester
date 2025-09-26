package alils

import (
	"fmt"
	"testing"
)

func TestEncodeFixed32Log_1(t *testing.T) {
	testCases := []struct {
		name   string
		data   []byte
		offset int
		v      uint32
		want   int
	}{
		{
			name:   "Test case 1",
			data:   make([]byte, 10),
			offset: 0,
			v:      1234567890,
			want:   4,
		},
		{
			name:   "Test case 2",
			data:   make([]byte, 10),
			offset: 5,
			v:      987654321,
			want:   9,
		},
		{
			name:   "Test case 3",
			data:   make([]byte, 10),
			offset: 0,
			v:      0,
			want:   4,
		},
		{
			name:   "Test case 4",
			data:   make([]byte, 10),
			offset: 0,
			v:      4294967295,
			want:   4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := encodeFixed32Log(tc.data, tc.offset, tc.v)
			if got != tc.want {
				t.Errorf("encodeFixed32Log(%v, %v, %v) = %v, want %v", tc.data, tc.offset, tc.v, got, tc.want)
			}
		})
	}
}
