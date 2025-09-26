package berror

import (
	"fmt"
	"testing"
)

func TestCodeBlock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lan := "go"
	code := "package main\n\nfunc main() {\n\tprintln(\"Hello, world!\")\n}"
	want := "```go\npackage main\n\nfunc main() {\n\tprintln(\"Hello, world!\")\n}\n```"
	got := codeBlock(lan, code)
	if got != want {
		t.Errorf("codeBlock(%q, %q) = %q; want %q", lan, code, got, want)
	}
}
