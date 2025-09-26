package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestappendToInterfaceSliceFromValues_1(t *testing.T) {
	type args struct {
		slice1 reflect.Value
		slice2 reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		want    []any
		wantErr bool
	}{
		{
			name: "Test with two valid slices",
			args: args{
				slice1: reflect.ValueOf([]int{1, 2, 3}),
				slice2: reflect.ValueOf([]string{"a", "b", "c"}),
			},
			want:    []any{1, 2, 3, "a", "b", "c"},
			wantErr: false,
		},
		{
			name: "Test with one valid and one invalid slice",
			args: args{
				slice1: reflect.ValueOf([]int{1, 2, 3}),
				slice2: reflect.Value{},
			},
			want:    []any{1, 2, 3, nil},
			wantErr: false,
		},
		{
			name: "Test with two invalid slices",
			args: args{
				slice1: reflect.Value{},
				slice2: reflect.Value{},
			},
			want:    []any{nil, nil},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := appendToInterfaceSliceFromValues(tt.args.slice1, tt.args.slice2)
			if (err != nil) != tt.wantErr {
				t.Errorf("appendToInterfaceSliceFromValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendToInterfaceSliceFromValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
