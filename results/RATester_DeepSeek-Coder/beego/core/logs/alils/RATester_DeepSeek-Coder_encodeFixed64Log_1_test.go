package alils

import (
	"fmt"
	"testing"
)

func TestEncodeFixed64Log_1(t *testing.T) {
	testCases := []struct {
		name   string
		data   []byte
		offset int
		v      uint64
		want   int
	}{
		{
			name:   "Test case 1",
			data:   make([]byte, 10),
			offset: 0,
			v:      0x0102030405060708,
			want:   8,
		},
		{
			name:   "Test case 2",
			data:   make([]byte, 10),
			offset: 0,
			v:      0x0807060504030201,
			want:   8,
		},
		{
			name:   "Test case 3",
			data:   make([]byte, 10),
			offset: 0,
			v:      0xffffffffffffffff,
			want:   8,
		},
		{
			name:   "Test case 4",
			data:   make([]byte, 10),
			offset: 0,
			v:      0x0000000000000000,
			want:   8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := encodeFixed64Log(tc.data, tc.offset, tc.v)
			if got != tc.want {
				t.Errorf("encodeFixed64Log() = %v, want %v", got, tc.want)
			}
		})
	}
}
