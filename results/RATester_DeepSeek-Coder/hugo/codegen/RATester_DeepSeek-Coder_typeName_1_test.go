package codegen

import (
	"fmt"
	"testing"
)

func TestTypeName_1(t *testing.T) {
	type args struct {
		name string
		pkg  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				name: "main.main",
				pkg:  "main",
			},
			want: "main",
		},
		{
			name: "Test case 2",
			args: args{
				name: "fmt.Println",
				pkg:  "fmt",
			},
			want: "Println",
		},
		{
			name: "Test case 3",
			args: args{
				name: "io.ReadAll",
				pkg:  "io",
			},
			want: "ReadAll",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := typeName(tt.args.name, tt.args.pkg); got != tt.want {
				t.Errorf("typeName() = %v, want %v", got, tt.want)
			}
		})
	}
}
