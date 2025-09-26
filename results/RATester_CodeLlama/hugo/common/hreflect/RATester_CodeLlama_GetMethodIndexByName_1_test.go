package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMethodIndexByName_1(t *testing.T) {
	type args struct {
		tp   reflect.Type
		name string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				tp:   reflect.TypeOf(struct{}{}),
				name: "Name",
			},
			want: -1,
		},
		{
			name: "case2",
			args: args{
				tp:   reflect.TypeOf(struct{ Name string }{}),
				name: "Name",
			},
			want: 0,
		},
		{
			name: "case3",
			args: args{
				tp:   reflect.TypeOf(struct{ Name string }{}),
				name: "Name1",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetMethodIndexByName(tt.args.tp, tt.args.name); got != tt.want {
				t.Errorf("GetMethodIndexByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
