package utils

import (
	"fmt"
	"testing"
)

func TestGet_19(t *testing.T) {
	type args struct {
		i     int
		args  []string
		input ArgString
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				i:     0,
				args:  []string{"default"},
				input: ArgString{"value1", "value2", "value3"},
			},
			want: "value1",
		},
		{
			name: "Test case 2",
			args: args{
				i:     2,
				args:  []string{"default"},
				input: ArgString{"value1", "value2", "value3"},
			},
			want: "value3",
		},
		{
			name: "Test case 3",
			args: args{
				i:     3,
				args:  []string{"default"},
				input: ArgString{"value1", "value2", "value3"},
			},
			want: "default",
		},
		{
			name: "Test case 4",
			args: args{
				i:     -1,
				args:  []string{"default"},
				input: ArgString{"value1", "value2", "value3"},
			},
			want: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.input.Get(tt.args.i, tt.args.args...); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
