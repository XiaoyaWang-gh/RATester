package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestGenerate_1(t *testing.T) {
	type args struct {
		value []byte
		kind  ast.NodeKind
	}
	tests := []struct {
		name string
		ids  *idFactory
		args args
		want []byte
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

			if got := tt.ids.Generate(tt.args.value, tt.args.kind); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("idFactory.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
