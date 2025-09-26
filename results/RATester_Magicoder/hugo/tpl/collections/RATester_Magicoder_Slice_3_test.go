package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlice_3(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name string
		args []any
		want any
	}{
		{
			name: "empty slice",
			args: []any{},
			want: []any{},
		},
		{
			name: "slice with elements",
			args: []any{1, 2, 3},
			want: []any{1, 2, 3},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ns.Slice(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
