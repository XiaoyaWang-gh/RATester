package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrimArr_1(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{
				arr: []string{"", "hello", "", "world"},
			},
			want: []string{"hello", "world"},
		},
		{
			name: "Test case 2",
			args: args{
				arr: []string{"", "", "", ""},
			},
			want: []string{},
		},
		{
			name: "Test case 3",
			args: args{
				arr: []string{"one", "two", "three"},
			},
			want: []string{"one", "two", "three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := TrimArr(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
