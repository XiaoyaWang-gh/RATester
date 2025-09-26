package codeblocks

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestGetLang_1(t *testing.T) {
	type args struct {
		node *ast.FencedCodeBlock
		src  []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getLang(tt.args.node, tt.args.src); got != tt.want {
				t.Errorf("getLang() = %v, want %v", got, tt.want)
			}
		})
	}
}
