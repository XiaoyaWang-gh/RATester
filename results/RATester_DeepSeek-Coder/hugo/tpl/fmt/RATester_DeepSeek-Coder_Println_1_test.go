package fmt

import (
	"fmt"
	"testing"
)

func TestPrintln_1(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestPrintln with string",
			args: args{args: []any{"test"}},
			want: "test\n",
		},
		{
			name: "TestPrintln with multiple arguments",
			args: args{args: []any{"test1", "test2"}},
			want: "test1 test2\n",
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
			if got := ns.Println(tt.args.args...); got != tt.want {
				t.Errorf("Namespace.Println() = %v, want %v", got, tt.want)
			}
		})
	}
}
