package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveArrVal_1(t *testing.T) {
	type args struct {
		arr []string
		val string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{
				arr: []string{"a", "b", "c", "d", "e"},
				val: "c",
			},
			want: []string{"a", "b", "d", "e"},
		},
		{
			name: "Test case 2",
			args: args{
				arr: []string{"a", "b", "c", "d", "e"},
				val: "f",
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "Test case 3",
			args: args{
				arr: []string{"a", "a", "b", "c", "a"},
				val: "a",
			},
			want: []string{"b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := RemoveArrVal(tt.args.arr, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveArrVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
