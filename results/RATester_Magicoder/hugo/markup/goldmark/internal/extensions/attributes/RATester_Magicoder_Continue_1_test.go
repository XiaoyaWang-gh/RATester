package attributes

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func TestContinue_1(t *testing.T) {
	tests := []struct {
		name   string
		a      *attrParser
		node   ast.Node
		reader text.Reader
		pc     parser.Context
		want   parser.State
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

			if got := tt.a.Continue(tt.node, tt.reader, tt.pc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("attrParser.Continue() = %v, want %v", got, tt.want)
			}
		})
	}
}
