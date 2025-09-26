package etag

import (
	"fmt"
	"reflect"
	"testing"
)

func TestappendUint_1(t *testing.T) {
	tests := []struct {
		name string
		dst  []byte
		n    uint32
		want []byte
	}{
		{
			name: "Test 1",
			dst:  []byte{},
			n:    123,
			want: []byte{'3', '2', '1'},
		},
		{
			name: "Test 2",
			dst:  []byte{},
			n:    0,
			want: []byte{'0'},
		},
		{
			name: "Test 3",
			dst:  []byte{},
			n:    1000,
			want: []byte{'0', '0', '0', '1'},
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
