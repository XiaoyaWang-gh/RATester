package parse

import (
	"fmt"
	"testing"
)

func TestType_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &elseNode{}
	if e.Type() != nodeElse {
		t.Errorf("e.Type() = %v, want %v", e.Type(), nodeElse)
	}
}
