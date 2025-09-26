package tcp

import (
	"fmt"
	"testing"
)

func TestExtend_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	chain1 := Chain{[]Constructor{}}
	chain2 := Chain{[]Constructor{}}

	extendedChain := chain1.Extend(chain2)

	if len(extendedChain.constructors) != 0 {
		t.Errorf("Expected empty constructors slice, got %v", extendedChain.constructors)
	}
}
