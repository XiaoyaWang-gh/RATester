package migration

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

	d := dataSlice{
		{created: 1},
		{created: 2},
		{created: 3},
	}
	i := 0
	j := 1
	if !d.Less(i, j) {
		t.Errorf("Less(%d, %d) = false, want true", i, j)
	}
}
