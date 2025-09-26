package cache

import (
	"fmt"
	"testing"
)

func TestGetInt64_1(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test with int",
			args: args{v: 1},
			want: 1,
		},
		{
			name: "Test with int32",
			args: args{v: int32(2)},
			want: 2,
		},
		{
			name: "Test with int64",
			args: args{v: int64(3)},
			want: 3,
		},
		{
			name: "Test with string",
			args: args{v: "4"},
			want: 4,
		},
		{
			name: "Test with other type",
			args: args{v: "5.5"},
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

			if got := GetInt64(tt.args.v); got != tt.want {
				t.Errorf("GetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
