package passthrough

import (
	"fmt"
	"testing"
)

func TestType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &passthroughContext{}
	p.typ = "inner"
	if p.Type() != "inner" {
		t.Errorf("Type() = %v, want %v", p.Type(), "inner")
	}
}
