package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFunction_1(t *testing.T) {
	type args struct {
		pc uintptr
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "testcase1",
			args: args{
				pc: 1,
			},
			want: []byte("testcase1"),
		},
		{
			name: "testcase2",
			args: args{
				pc: 2,
			},
			want: []byte("testcase2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := function(tt.args.pc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("function() = %v, want %v", got, tt.want)
			}
		})
	}
}
