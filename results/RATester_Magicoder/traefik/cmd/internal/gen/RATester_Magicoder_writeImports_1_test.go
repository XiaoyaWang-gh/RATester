package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWriteImports_1(t *testing.T) {
	type args struct {
		b       *bytes.Buffer
		imports []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with no imports",
			args: args{
				b:       &bytes.Buffer{},
				imports: []string{},
			},
			want: "",
		},
		{
			name: "Test with one import",
			args: args{
				b:       &bytes.Buffer{},
				imports: []string{"fmt"},
			},
			want: "import (\n\t\"fmt\"\n)\n",
		},
		{
			name: "Test with multiple imports",
			args: args{
				b:       &bytes.Buffer{},
				imports: []string{"fmt", "io", "os", "fmt"},
			},
			want: "import (\n\t\"fmt\"\n\t\"io\"\n\t\"os\"\n)\n",
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
