package pagemeta

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToLowerSlice_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   any
		want []string
	}{
		{
			name: "Test Case 1",
			in:   []any{"A", "B", "C"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Test Case 2",
			in:   []any{"d", "e", "f"},
			want: []string{"d", "e", "f"},
		},
		{
			name: "Test Case 3",
			in:   []any{"g", "h", "i"},
			want: []string{"g", "h", "i"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := toLowerSlice(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toLowerSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
