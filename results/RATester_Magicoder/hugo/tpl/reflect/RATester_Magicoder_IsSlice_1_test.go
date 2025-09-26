package reflect

import (
	"fmt"
	"testing"
)

func TestIsSlice_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name string
		v    any
		want bool
	}{
		{
			name: "slice",
			v:    []int{1, 2, 3},
			want: true,
		},
		{
			name: "not slice",
			v:    "not slice",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ns.IsSlice(tt.v); got != tt.want {
				t.Errorf("IsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
