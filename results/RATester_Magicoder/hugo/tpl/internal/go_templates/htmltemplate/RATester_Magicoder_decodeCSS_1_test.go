package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestdecodeCSS_1(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{
			name: "test case 1",
			s:    []byte("\\000"),
			want: []byte{0},
		},
		{
			name: "test case 2",
			s:    []byte("\\000\\001"),
			want: []byte{0, 1},
		},
		{
			name: "test case 3",
			s:    []byte("\\000\\001\\002"),
			want: []byte{0, 1, 2},
		},
		{
			name: "test case 4",
			s:    []byte("\\000\\001\\002\\003"),
			want: []byte{0, 1, 2, 3},
		},
		{
			name: "test case 5",
			s:    []byte("\\000\\001\\002\\003\\004"),
			want: []byte{0, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := decodeCSS(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeCSS() = %v, want %v", got, tt.want)
			}
		})
	}
}
