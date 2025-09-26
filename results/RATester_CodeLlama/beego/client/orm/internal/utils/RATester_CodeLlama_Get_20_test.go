package utils

import (
	"fmt"
	"testing"
)

func TestGet_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := ArgInt{1, 2, 3}
	i := 1
	args := []int{4, 5, 6}
	r := a.Get(i, args...)
	if r != 2 {
		t.Errorf("a.Get(1, 4, 5, 6) = %d; want 2", r)
	}
}
