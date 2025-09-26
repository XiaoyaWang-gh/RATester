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

	c := Chain{}
	constructors := []Constructor{
		func(h Handler) (Handler, error) {
			return h, nil
		},
		func(h Handler) (Handler, error) {
			return h, nil
		},
	}
	c = c.Append(constructors...)
	if len(c.constructors) != 2 {
		t.Errorf("Expected 2 constructors, got %d", len(c.constructors))
	}
}
