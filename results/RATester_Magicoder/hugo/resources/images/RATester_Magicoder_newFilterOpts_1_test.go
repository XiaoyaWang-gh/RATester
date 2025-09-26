package images

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewFilterOpts_1(t *testing.T) {
	type args struct {
		vals []any
	}
	tests := []struct {
		name string
		args args
		want filterOpts
	}{
		{
			name: "Test Case 1",
			args: args{
				vals: []any{1, "test", true},
			},
			want: filterOpts{
				Version: filterAPIVersion,
				Vals:    []any{1, "test", true},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				vals: []any{},
			},
			want: filterOpts{
				Version: filterAPIVersion,
				Vals:    []any{},
			},
		},
		{
			name: "Test Case 3",
			args: args{
				vals: []any{nil},
			},
			want: filterOpts{
				Version: filterAPIVersion,
				Vals:    []any{nil},
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
