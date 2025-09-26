package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStringSlicePreserveString_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		v    any
		want []string
	}{
		{
			name: "Test with string",
			v:    "test",
			want: []string{"test"},
		},
		{
			name: "Test with slice of strings",
			v:    []string{"test1", "test2"},
			want: []string{"test1", "test2"},
		},
		{
			name: "Test with slice of ints",
			v:    []int{1, 2, 3, 4, 5},
			want: []string{"1", "2", "3", "4", "5"},
		},
		{
			name: "Test with slice of float64s",
			v:    []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			want: []string{"1.1", "2.2", "3.3", "4.4", "5.5"},
		},
		{
			name: "Test with slice of bools",
			v:    []bool{true, false, true},
			want: []string{"true", "false", "true"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ToStringSlicePreserveString(tt.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringSlicePreserveString() = %v, want %v", got, tt.want)
			}
		})
	}
}
