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

	chain1 := Chain{
		constructors: []Constructor{
			func(h Handler) (Handler, error) {
				return nil, nil
			},
		},
	}

	chain2 := Chain{
		constructors: []Constructor{
			func(h Handler) (Handler, error) {
				return nil, nil
			},
		},
	}

	extendedChain := chain1.Extend(chain2)

	if len(extendedChain.constructors) != 2 {
		t.Errorf("Expected length of constructors to be 2, got %d", len(extendedChain.constructors))
	}
}
