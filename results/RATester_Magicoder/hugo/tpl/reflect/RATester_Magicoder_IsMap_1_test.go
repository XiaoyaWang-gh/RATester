package reflect

import (
	"fmt"
	"testing"
)

func TestIsMap_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name string
		v    any
		want bool
	}{
		{
			name: "Test with map",
			v:    map[string]int{"one": 1, "two": 2},
			want: true,
		},
		{
			name: "Test with slice",
			v:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "Test with nil",
			v:    nil,
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

			if got := ns.IsMap(tt.v); got != tt.want {
				t.Errorf("IsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
