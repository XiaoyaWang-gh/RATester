package codegen

import (
	"fmt"
	"testing"
)

func TestVarName_1(t *testing.T) {
	varName := func(name string) string {
		name = firstToLower(name)

		switch name {
		case "type":
			name = "typ"
		case "package":
			name = "pkg"

		case "len":
			name = "length"
		}

		return name
	}

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				name: "Type",
			},
			want: "typ",
		},
		{
			name: "case 2",
			args: args{
				name: "Package",
			},
			want: "pkg",
		},
		{
			name: "case 3",
			args: args{
				name: "Len",
			},
			want: "length",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := varName(tt.args.name); got != tt.want {
				t.Errorf("varName() = %v, want %v", got, tt.want)
			}
		})
	}
}
