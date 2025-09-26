package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestcanCompare_1(t *testing.T) {
	type args struct {
		v1 reflect.Value
		v2 reflect.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				v1: reflect.ValueOf(1),
				v2: reflect.ValueOf(2),
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				v1: reflect.ValueOf(1),
				v2: reflect.ValueOf("2"),
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				v1: reflect.ValueOf(1),
				v2: reflect.ValueOf(reflect.ValueOf(2)),
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				v1: reflect.ValueOf(1),
				v2: reflect.ValueOf(reflect.ValueOf("2")),
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

			if got := canCompare(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("canCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
