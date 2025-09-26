package glob

import (
	"fmt"
	"testing"
)

func TestResolveRootDir_1(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "/a/b/c", want: "a"},
		{input: "/a/b/c/d", want: "a/b"},
		{input: "/a/b/c/d/e", want: "a/b/c"},
		{input: "/a/b/c/d/e/f", want: "a/b/c/d"},
		{input: "/a/b/c/d/e/f/g", want: "a/b/c/d/e"},
		{input: "/a", want: ""},
		{input: "/", want: ""},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("input=%s", tt.input), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := ResolveRootDir(tt.input)
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
