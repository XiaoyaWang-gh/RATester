package parse

import (
	"fmt"
	"testing"
)

func TestSetPos_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := &IdentifierNode{Ident: "foo"}
	pos := Pos(10)
	i.SetPos(pos)
	if i.Pos != pos {
		t.Errorf("expected %d, got %d", pos, i.Pos)
	}
}
