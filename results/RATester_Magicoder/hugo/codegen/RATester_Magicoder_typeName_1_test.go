package codegen

import (
	"fmt"
	"testing"
)

func TesttypeName_1(t *testing.T) {
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
				name: "pkg.Name",
				pkg:  "pkg",
			},
			want: "Name",
		},
		{
			name: "Test case 2",
			args: args{
				name: "pkg.Name",
				pkg:  "pkg",
			},
			want: "Name",
		},
		// Add more test cases as needed
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
