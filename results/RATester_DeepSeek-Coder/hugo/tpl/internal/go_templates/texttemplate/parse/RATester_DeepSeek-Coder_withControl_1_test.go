package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithControl_1(t *testing.T) {
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
		want   Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1 := &Tree{
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
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := t1.withControl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.withControl() = %v, want %v", got, tt.want)
			}
		})
	}
}
