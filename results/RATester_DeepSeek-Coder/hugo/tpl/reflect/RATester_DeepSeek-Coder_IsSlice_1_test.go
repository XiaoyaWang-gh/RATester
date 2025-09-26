package reflect

import (
	"fmt"
	"testing"
)

func TestIsSlice_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	tests := []struct {
		name string
		v    any
		want bool
	}{
		{
			name: "Test with a slice",
			v:    []int{1, 2, 3},
			want: true,
		},
		{
			name: "Test with a string",
			v:    "test",
			want: false,
		},
		{
			name: "Test with a map",
			v:    map[string]int{"one": 1, "two": 2, "three": 3},
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
				t.Errorf("Namespace.IsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
