package fmt

import (
	"fmt"
	"testing"
)

func TestPrintf_1(t *testing.T) {
	type args struct {
		format string
		args   []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestPrintf_0",
			args: args{
				format: "Hello, %s",
				args:   []any{"World"},
			},
			want: "Hello, World",
		},
		{
			name: "TestPrintf_1",
			args: args{
				format: "The value is %d",
				args:   []any{42},
			},
			want: "The value is 42",
		},
		{
			name: "TestPrintf_2",
			args: args{
				format: "The value is %.2f",
				args:   []any{3.14159},
			},
			want: "The value is 3.14",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ns := &Namespace{}
			if got := ns.Printf(tt.args.format, tt.args.args...); got != tt.want {
				t.Errorf("Namespace.Printf() = %v, want %v", got, tt.want)
			}
		})
	}
}
