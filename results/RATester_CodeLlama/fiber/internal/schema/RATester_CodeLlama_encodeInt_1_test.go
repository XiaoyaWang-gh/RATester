package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeInt_1(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				v: reflect.ValueOf(1),
			},
			want: "1",
		},
		{
			name: "case 2",
			args: args{
				v: reflect.ValueOf(1.1),
			},
			want: "1",
		},
		{
			name: "case 3",
			args: args{
				v: reflect.ValueOf("1"),
			},
			want: "1",
		},
		{
			name: "case 4",
			args: args{
				v: reflect.ValueOf([]int{1}),
			},
			want: "1",
		},
		{
			name: "case 5",
			args: args{
				v: reflect.ValueOf(map[string]int{"1": 1}),
			},
			want: "1",
		},
		{
			name: "case 6",
			args: args{
				v: reflect.ValueOf(struct {
					Name string
				}{Name: "1"}),
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := encodeInt(tt.args.v); got != tt.want {
				t.Errorf("encodeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
