package web

import (
	"fmt"
	"testing"
)

func TestLess_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := ControllerCommentsSlice{
		{Router: "a"},
		{Router: "b"},
		{Router: "c"},
	}
	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			if p.Less(i, j) {
				t.Errorf("Less(%d, %d) = true", i, j)
			}
		}
	}
}
