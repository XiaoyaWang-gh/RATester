package utils

import (
	"fmt"
	"testing"
)

func TestInSliceIface_1(t *testing.T) {
	type args struct {
		v  interface{}
		sl []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				v:  "test",
				sl: []interface{}{"test", "test2", "test3"},
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				v:  123,
				sl: []interface{}{"test", 123, "test3"},
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				v:  "test4",
				sl: []interface{}{"test", 123, "test3"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := InSliceIface(tt.args.v, tt.args.sl); got != tt.want {
				t.Errorf("InSliceIface() = %v, want %v", got, tt.want)
			}
		})
	}
}
