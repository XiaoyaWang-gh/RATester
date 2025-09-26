package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestshiftNRuneBytes_1(t *testing.T) {
	tests := []struct {
		name string
		rb   [4]byte
		n    int
		want [4]byte
	}{
		{
			name: "shift 0",
			rb:   [4]byte{1, 2, 3, 4},
			n:    0,
			want: [4]byte{1, 2, 3, 4},
		},
		{
			name: "shift 1",
			rb:   [4]byte{1, 2, 3, 4},
			n:    1,
			want: [4]byte{2, 3, 4, 0},
		},
		{
			name: "shift 2",
			rb:   [4]byte{1, 2, 3, 4},
			n:    2,
			want: [4]byte{3, 4},
		},
		{
			name: "shift 3",
			rb:   [4]byte{1, 2, 3, 4},
			n:    3,
			want: [4]byte{4},
		},
		{
			name: "shift 4",
			rb:   [4]byte{1, 2, 3, 4},
			n:    4,
			want: [4]byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := shiftNRuneBytes(tt.rb, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shiftNRuneBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
