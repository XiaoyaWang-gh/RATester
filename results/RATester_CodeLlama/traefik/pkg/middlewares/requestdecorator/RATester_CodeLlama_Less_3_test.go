package requestdecorator

import (
	"fmt"
	"testing"
)

func TestLess_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := byTTL{
		{TTL: 1},
		{TTL: 2},
		{TTL: 3},
	}
	i := 0
	j := 1
	if !a.Less(i, j) {
		t.Errorf("Less(%d, %d) = false, want true", i, j)
	}
	if a.Less(j, i) {
		t.Errorf("Less(%d, %d) = true, want false", j, i)
	}
}
