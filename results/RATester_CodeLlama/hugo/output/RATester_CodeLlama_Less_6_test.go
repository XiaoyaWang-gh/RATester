package output

import (
	"fmt"
	"testing"
)

func TestLess_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	formats := Formats{
		{Name: "a", Weight: 1},
		{Name: "b", Weight: 2},
		{Name: "c", Weight: 3},
		{Name: "d", Weight: 4},
		{Name: "e", Weight: 5},
	}

	for i := 0; i < len(formats)-1; i++ {
		for j := i + 1; j < len(formats); j++ {
			if formats.Less(i, j) != (formats[i].Weight < formats[j].Weight) {
				t.Errorf("Less(%d, %d) = %t, want %t", i, j, formats.Less(i, j), formats[i].Weight < formats[j].Weight)
			}
		}
	}
}
