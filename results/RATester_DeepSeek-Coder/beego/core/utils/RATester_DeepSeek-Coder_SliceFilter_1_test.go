package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceFilter_1(t *testing.T) {
	type args struct {
		slice []interface{}
		a     filtertype
	}
	tests := []struct {
		name        string
		args        args
		wantFtslice []interface{}
	}{
		{
			name: "Test case 1",
			args: args{
				slice: []interface{}{1, 2, 3, 4, 5},
				a: func(i interface{}) bool {
					return i.(int)%2 == 0
				},
			},
			wantFtslice: []interface{}{2, 4},
		},
		{
			name: "Test case 2",
			args: args{
				slice: []interface{}{"apple", "banana", "cherry", "date"},
				a: func(i interface{}) bool {
					_, ok := i.(string)
					return ok
				},
			},
			wantFtslice: []interface{}{"apple", "banana", "cherry", "date"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotFtslice := SliceFilter(tt.args.slice, tt.args.a); !reflect.DeepEqual(gotFtslice, tt.wantFtslice) {
				t.Errorf("SliceFilter() = %v, want %v", gotFtslice, tt.wantFtslice)
			}
		})
	}
}
