package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestWriteImports_1(t *testing.T) {
	type args struct {
		b       *strings.Builder
		imports []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty imports",
			args: args{
				b:       &strings.Builder{},
				imports: []string{},
			},
			want: "",
		},
		{
			name: "single import",
			args: args{
				b:       &strings.Builder{},
				imports: []string{"fmt"},
			},
			want: "import (\n\t\"fmt\"\n)\n",
		},
		{
			name: "multiple imports",
			args: args{
				b:       &strings.Builder{},
				imports: []string{"fmt", "os", "fmt"},
			},
			want: "import (\n\t\"fmt\"\n\t\"os\"\n)\n",
		},
		{
			name: "sorted imports",
			args: args{
				b:       &strings.Builder{},
				imports: []string{"fmt", "os"},
			},
			want: "import (\n\t\"fmt\"\n\t\"os\"\n)\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := fileWriter{}
			f.writeImports(tt.args.b, tt.args.imports)
			if got := tt.args.b.String(); got != tt.want {
				t.Errorf("writeImports() = %v, want %v", got, tt.want)
			}
		})
	}
}
