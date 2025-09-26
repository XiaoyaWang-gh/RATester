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
				arr: []string{"apple", "banana", "cherry"},
				val: "banana",
			},
			want: []string{"apple", "cherry"},
		},
		{
			name: "Test case 2",
			args: args{
				arr: []string{"apple", "banana", "cherry"},
				val: "orange",
			},
			want: []string{"apple", "banana", "cherry"},
		},
		{
			name: "Test case 3",
			args: args{
				arr: []string{"apple", "banana", "cherry"},
				val: "apple",
			},
			want: []string{"banana", "cherry"},
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
