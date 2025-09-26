package fmt

import (
	"fmt"
	"testing"
)

func TestPrint_1(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestPrint_0",
			args: args{
				args: []any{1, "test", true},
			},
			want: "1 test true",
		},
		{
			name: "TestPrint_1",
			args: args{
				args: []any{},
			},
			want: "",
		},
		{
			name: "TestPrint_2",
			args: args{
				args: []any{1},
			},
			want: "1",
		},
		{
			name: "TestPrint_3",
			args: args{
				args: []any{"test", true},
			},
			want: "test true",
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
			if got := ns.Print(tt.args.args...); got != tt.want {
				t.Errorf("Namespace.Print() = %v, want %v", got, tt.want)
			}
		})
	}
}
