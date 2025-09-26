package parse

import (
	"fmt"
	"testing"
)

func TestStartParse_1(t *testing.T) {
	type args struct {
		funcs   []map[string]any
		lex     *lexer
		treeSet map[string]*Tree
	}
	tests := []struct {
		name string
		t    *Tree
		args args
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

			tt.t.startParse(tt.args.funcs, tt.args.lex, tt.args.treeSet)
		})
	}
}
