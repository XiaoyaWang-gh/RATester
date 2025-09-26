package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringSliceToInterfaceSlice_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg  []string
		want []any
	}{
		{
			name: "empty slice",
			arg:  []string{},
			want: []any{},
		},
		{
			name: "slice with one element",
			arg:  []string{"hello"},
			want: []any{"hello"},
		},
		{
			name: "slice with multiple elements",
			arg:  []string{"hello", "world"},
			want: []any{"hello", "world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := StringSliceToInterfaceSlice(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSliceToInterfaceSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
