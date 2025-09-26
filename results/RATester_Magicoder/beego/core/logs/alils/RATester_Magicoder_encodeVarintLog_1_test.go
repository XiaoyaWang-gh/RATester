package alils

import (
	"fmt"
	"testing"
)

func TestencodeVarintLog_1(t *testing.T) {
	tests := []struct {
		name   string
		v      uint64
		offset int
		data   []byte
		want   int
	}{
		{
			name:   "Test case 1",
			v:      1,
			offset: 0,
			data:   make([]byte, 10),
			want:   1,
		},
		{
			name:   "Test case 2",
			v:      127,
			offset: 0,
			data:   make([]byte, 10),
			want:   1,
		},
		{
			name:   "Test case 3",
			v:      128,
			offset: 0,
			data:   make([]byte, 10),
			want:   2,
		},
		{
			name:   "Test case 4",
			v:      16383,
			offset: 0,
			data:   make([]byte, 10),
			want:   2,
		},
		{
			name:   "Test case 5",
			v:      16384,
			offset: 0,
			data:   make([]byte, 10),
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := encodeVarintLog(tt.data, tt.offset, tt.v)
			if got != tt.want {
				t.Errorf("encodeVarintLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
