package main

import (
	"fmt"
	"testing"
)

func TestcommonReplace_1(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    string
	}{
		{
			name:    "test.go",
			content: "package template\n",
			want: `// +build go1.13,!windows

package template
`,
		},
		{
			name:    "test_test.go",
			content: "package template_test\n",
			want: `// +build go1.13

package template_test
`,
		},
		{
			name:    "parse.go",
			content: "package parse\n",
			want: `// +build go1.13

package parse
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := commonReplace(tt.name, tt.content); got != tt.want {
				t.Errorf("commonReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}
