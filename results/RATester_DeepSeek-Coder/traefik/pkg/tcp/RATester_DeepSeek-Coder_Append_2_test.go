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
	chain2 := chain1.Append(func(h Handler) (Handler, error) {
		return h, nil
	})

	if len(chain1.constructors) != 0 {
		t.Errorf("Expected chain1 to have 0 constructors, got %d", len(chain1.constructors))
	}

	if len(chain2.constructors) != 1 {
		t.Errorf("Expected chain2 to have 1 constructor, got %d", len(chain2.constructors))
	}

	chain3 := chain2.Append(func(h Handler) (Handler, error) {
		return h, nil
	}, func(h Handler) (Handler, error) {
		return h, nil
	})

	if len(chain3.constructors) != 3 {
		t.Errorf("Expected chain3 to have 3 constructors, got %d", len(chain3.constructors))
	}
}
