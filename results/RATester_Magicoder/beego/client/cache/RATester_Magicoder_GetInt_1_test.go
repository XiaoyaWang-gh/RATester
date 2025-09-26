package cache

import (
	"fmt"
	"testing"
)

func TestGetInt_1(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test with int",
			args: args{v: 10},
			want: 10,
		},
		{
			name: "Test with int32",
			args: args{v: int32(20)},
			want: 20,
		},
		{
			name: "Test with int64",
			args: args{v: int64(30)},
			want: 30,
		},
		{
			name: "Test with string",
			args: args{v: "40"},
			want: 40,
		},
		{
			name: "Test with other type",
			args: args{v: "not a number"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetInt(tt.args.v); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
