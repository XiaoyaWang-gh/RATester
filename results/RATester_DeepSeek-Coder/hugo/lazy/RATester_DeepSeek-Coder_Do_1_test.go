package lazy

import (
	"fmt"
	"testing"
)

func TestDo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &onceMore{}

	called := false
	f := func() { called = true }

	o.Do(f)
	if !called {
		t.Errorf("expected function to be called once")
	}

	called = false
	o.Do(f)
	if called {
		t.Errorf("expected function not to be called again")
	}
}
