package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func Testfomateinfo_1(t *testing.T) {
	type args struct {
		headlen int
		data    []interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test Case 1",
			args: args{
				headlen: 1,
				data:    []interface{}{"test1", "test2"},
			},
			want: []byte("    [test1, test2]\n"),
		},
		{
			name: "Test Case 2",
			args: args{
				headlen: 1,
				data:    []interface{}{"test3", "test4"},
			},
			want: []byte("    [test3, test4]\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := fomateinfo(tt.args.headlen, tt.args.data...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fomateinfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
