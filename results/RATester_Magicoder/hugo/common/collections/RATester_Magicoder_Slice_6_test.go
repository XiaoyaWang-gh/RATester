package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlice_6(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Test case 1",
			args: args{
				args: []any{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test case 2",
			args: args{
				args: []any{"a", "b", "c", "d", "e"},
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "Test case 3",
			args: args{
				args: []any{true, false, true, false, true},
			},
			want: []bool{true, false, true, false, true},
		},
		{
			name: "Test case 4",
			args: args{
				args: []any{1.1, 2.2, 3.3, 4.4, 5.5},
			},
			want: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
		},
		{
			name: "Test case 5",
			args: args{
				args: []any{nil, nil, nil, nil, nil},
			},
			want: []any{nil, nil, nil, nil, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Slice(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
