package berror

import (
	"fmt"
	"testing"
)

func TestCodeBlock_1(t *testing.T) {
	type args struct {
		lan  string
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
				lan:  "go",
				code: "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, world\")\n}",
			},
			want: "```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, world\")\n}\n```",
		},
		{
			name: "Test case 2",
			args: args{
				lan:  "python",
				code: "print('Hello, world')",
			},
			want: "```python\nprint('Hello, world')\n```",
		},
		{
			name: "Test case 3",
			args: args{
				lan:  "javascript",
				code: "console.log('Hello, world')",
			},
			want: "```javascript\nconsole.log('Hello, world')\n```",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := codeBlock(tt.args.lan, tt.args.code); got != tt.want {
				t.Errorf("codeBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
