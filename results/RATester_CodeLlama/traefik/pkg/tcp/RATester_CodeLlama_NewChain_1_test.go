package tcp

import (
	"fmt"
	"testing"
)

func TestNewChain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	constructors := []Constructor{
		func(h Handler) (Handler, error) {
			return h, nil
		},
		func(h Handler) (Handler, error) {
			return h, nil
		},
	}
	chain := NewChain(constructors...)
	if len(chain.constructors) != len(constructors) {
		t.Errorf("NewChain() = %v, want %v", len(chain.constructors), len(constructors))
	}
}
