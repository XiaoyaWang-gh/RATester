package parse

import (
	"fmt"
	"testing"
)

func TestBackup_1(t *testing.T) {
	type fields struct {
		Name      string
		ParseName string
		Root      *ListNode
		Mode      Mode
		text      string
		funcs     []map[string]any
		lex       *lexer
		token     [3]item
		peekCount int
		vars      []string
		treeSet   map[string]*Tree
	}
	tests := []struct {
		name   string
		fields fields
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

			tr := &Tree{
				Name:      tt.fields.Name,
				ParseName: tt.fields.ParseName,
				Root:      tt.fields.Root,
				Mode:      tt.fields.Mode,
				text:      tt.fields.text,
				funcs:     tt.fields.funcs,
				lex:       tt.fields.lex,
				token:     tt.fields.token,
				peekCount: tt.fields.peekCount,
				vars:      tt.fields.vars,
				treeSet:   tt.fields.treeSet,
			}
			tr.backup()
		})
	}
}
