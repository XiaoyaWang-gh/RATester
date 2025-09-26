package types

import (
	"fmt"
	"testing"
)

func TestToString_5(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				v: "Hello, World!",
			},
			want: "Hello, World!",
		},
		{
			name: "Test case 2",
			args: args{
				v: 123,
			},
			want: "123",
		},
		{
			name: "Test case 3",
			args: args{
				v: true,
			},
			want: "true",
		},
		{
			name: "Test case 4",
			args: args{
				v: nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ToString(tt.args.v); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
