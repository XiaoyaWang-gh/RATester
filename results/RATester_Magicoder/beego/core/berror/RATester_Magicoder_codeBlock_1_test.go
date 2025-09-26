package berror

import (
	"fmt"
	"testing"
)

func TestcodeBlock_1(t *testing.T) {
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
				code: "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")",
			},
			want: "```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")",
		},
		{
			name: "Test case 2",
			args: args{
				lan:  "python",
				code: "print(\"Hello, World!\")",
			},
			want: "```python\nprint(\"Hello, World!\")",
		},
		{
			name: "Test case 3",
			args: args{
				lan:  "java",
				code: "public class HelloWorld {\n\tpublic static void main(String[] args) {\n\t\tSystem.out.println(\"Hello, World!\");\n\t}\n}",
			},
			want: "```java\npublic class HelloWorld {\n\tpublic static void main(String[] args) {\n\t\tSystem.out.println(\"Hello, World!\");\n\t}\n}",
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
