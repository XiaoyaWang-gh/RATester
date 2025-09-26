package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestGenerate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ids := &idFactory{
		idType: "id",
		vals:   map[string]struct{}{},
	}
	value := []byte("value")
	kind := ast.KindHeading
	want := []byte("heading")
	got := ids.Generate(value, kind)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ids.Generate(value, kind) = %v, want %v", got, want)
	}
}
