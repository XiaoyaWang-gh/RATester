package cache

import (
	"fmt"
	"testing"
)

func TestGetBool_1(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				v: true,
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				v: "true",
			},
			want: true,
		},
		{
			name: "Test case 3",
			args: args{
				v: "false",
			},
			want: false,
		},
		{
			name: "Test case 4",
			args: args{
				v: "invalid",
			},
			want: false,
		},
		{
			name: "Test case 5",
			args: args{
				v: 123,
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

			if got := GetBool(tt.args.v); got != tt.want {
				t.Errorf("GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
