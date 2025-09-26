package template

import (
	"fmt"
	"testing"
)

func TestIsInScriptLiteral_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var s state
	if isInScriptLiteral(s) {
		t.Errorf("isInScriptLiteral(%v) = true, want false", s)
	}
}
