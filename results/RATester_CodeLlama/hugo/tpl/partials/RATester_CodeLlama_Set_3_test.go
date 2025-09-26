package partials

import (
	"fmt"
	"testing"
)

func TestSet_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &contextWrapper{}
	in := "hello"
	c.Set(in)
	if c.Result != in {
		t.Errorf("c.Result = %v, want %v", c.Result, in)
	}
}
