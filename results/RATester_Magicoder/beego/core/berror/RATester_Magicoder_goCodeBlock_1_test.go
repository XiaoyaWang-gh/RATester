package berror

import (
	"fmt"
	"testing"
)

func TestgoCodeBlock_1(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				code: "fmt.Println(\"Hello, World!\")",
			},
			want: "```go\nfmt.Println(\"Hello, World!\")",
		},
		{
			name: "Test case 2",
			args: args{
				code: "func main() {\n\tfmt.Println(\"Hello, World!\")",
			},
			want: "```go\nfunc main() {\n\tfmt.Println(\"Hello, World!\")",
		},
		{
			name: "Test case 3",
			args: args{
				code: "func main() {\n\tfmt.Println(\"Hello, World!\")",
			},
			want: "```go\nfunc main() {\n\tfmt.Println(\"Hello, World!\")",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := goCodeBlock(tt.args.code); got != tt.want {
				t.Errorf("goCodeBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
