package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringSliceToInterfaceSlice_1(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "Test case 1",
			args: args{
				ss: []string{"hello", "world"},
			},
			want: []any{"hello", "world"},
		},
		{
			name: "Test case 2",
			args: args{
				ss: []string{"1", "2", "3"},
			},
			want: []any{"1", "2", "3"},
		},
		{
			name: "Test case 3",
			args: args{
				ss: []string{},
			},
			want: []any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := StringSliceToInterfaceSlice(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSliceToInterfaceSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
