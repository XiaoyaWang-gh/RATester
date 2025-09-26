package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceReduce_1(t *testing.T) {
	type args struct {
		slice []interface{}
		a     reducetype
	}
	tests := []struct {
		name       string
		args       args
		wantDslice []interface{}
	}{
		{
			name: "Test case 1",
			args: args{
				slice: []interface{}{1, 2, 3, 4, 5},
				a: func(i interface{}) interface{} {
					return i.(int) * 2
				},
			},
			wantDslice: []interface{}{2, 4, 6, 8, 10},
		},
		{
			name: "Test case 2",
			args: args{
				slice: []interface{}{"a", "b", "c", "d", "e"},
				a: func(i interface{}) interface{} {
					return i.(string) + "x"
				},
			},
			wantDslice: []interface{}{"ax", "bx", "cx", "dx", "ex"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotDslice := SliceReduce(tt.args.slice, tt.args.a); !reflect.DeepEqual(gotDslice, tt.wantDslice) {
				t.Errorf("SliceReduce() = %v, want %v", gotDslice, tt.wantDslice)
			}
		})
	}
}
