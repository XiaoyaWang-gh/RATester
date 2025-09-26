package attributes

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestNew_38(t *testing.T) {
	type args struct {
		astAttributes []ast.Attribute
		ownerType     AttributesOwnerType
	}
	tests := []struct {
		name string
		args args
		want *AttributesHolder
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

			got := New(tt.args.astAttributes, tt.args.ownerType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
