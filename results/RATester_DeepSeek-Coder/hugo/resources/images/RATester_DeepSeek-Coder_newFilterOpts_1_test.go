package images

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFilterOpts_1(t *testing.T) {
	type args struct {
		vals []any
	}
	tests := []struct {
		name string
		args args
		want filterOpts
	}{
		{
			name: "Test with no arguments",
			args: args{
				vals: []any{},
			},
			want: filterOpts{
				Version: filterAPIVersion,
				Vals:    []any{},
			},
		},
		{
			name: "Test with one argument",
			args: args{
				vals: []any{1},
			},
			want: filterOpts{
				Version: filterAPIVersion,
				Vals:    []any{1},
			},
		},
		{
			name: "Test with multiple arguments",
			args: args{
				vals: []any{1, "two", 3.0},
			},
			want: filterOpts{
				Version: filterAPIVersion,
				Vals:    []any{1, "two", 3.0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newFilterOpts(tt.args.vals...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFilterOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}
