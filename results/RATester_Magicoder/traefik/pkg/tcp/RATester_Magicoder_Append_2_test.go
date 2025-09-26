package tcp

import (
	"fmt"
	"testing"
)

func TestAppend_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	chain1 := Chain{[]Constructor{}}
	chain2 := Chain{[]Constructor{}}

	chain1 = chain1.Append(func(h Handler) (Handler, error) {
		return h, nil
	})

	chain2 = chain2.Append(func(h Handler) (Handler, error) {
		return h, nil
	}, func(h Handler) (Handler, error) {
		return h, nil
	})

	if len(chain1.constructors) != 1 {
		t.Errorf("Expected chain1 to have 1 constructor, got %d", len(chain1.constructors))
	}

	if len(chain2.constructors) != 2 {
		t.Errorf("Expected chain2 to have 2 constructors, got %d", len(chain2.constructors))
	}
}
