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
			name: "test_newFilterOpts_0",
			args: args{
				vals: []any{
					"test_newFilterOpts_0",
				},
			},
			want: filterOpts{
				Version: filterAPIVersion,
				Vals:    []any{"test_newFilterOpts_0"},
			},
		},
		{
			name: "test_newFilterOpts_1",
			args: args{
				vals: []any{
					"test_newFilterOpts_1",
					"test_newFilterOpts_2",
				},
			},
			want: filterOpts{
				Version: filterAPIVersion,
				Vals:    []any{"test_newFilterOpts_1", "test_newFilterOpts_2"},
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
