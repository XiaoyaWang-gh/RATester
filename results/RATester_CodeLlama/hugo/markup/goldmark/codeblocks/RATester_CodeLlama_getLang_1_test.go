package codeblocks

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestGetLang_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	node := &ast.FencedCodeBlock{}
	src := []byte("")
	lang := getLang(node, src)
	if lang != "" {
		t.Errorf("getLang() = %v, want %v", lang, "")
	}
}
